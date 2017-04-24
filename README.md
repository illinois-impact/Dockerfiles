

## `binfmt_misc` register

Register `qemu-*-static` for all supported processors except the current one

* `docker run --rm --privileged multiarch/qemu-user-static:register`

Same as above, but remove all registered `binfmt_misc` before

* `docker run --rm --privileged multiarch/qemu-user-static:register --reset`
