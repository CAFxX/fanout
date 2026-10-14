package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]int)
    var v3 uint8 = 2
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 uint64 = v1
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(fi6, fi6, fi6)
            v7 := []int{}
            var v8 bool = false
            _ = fi6
            _ = v7
            _ = v8
        }
        v1 = (v5 ^ v5)
        _ = fi4
        _ = v5
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v1 -= v1
        _ = fi9
    }
    if true {
        for fi10 := 0; fi10 < 10; fi10++ {
            v1 -= v1
            if (-286 == 518) {
                var v11 uint8 = uint8(v1)
                var v12 int16 = 349
                _ = v11
                _ = v12
            } else {
                if true {
                    delete(v2, 427)
                    if (true && ((!(true || true)) || false)) {
                        var v13 uint = 65535
                        switch fi10 {
                        case -585:
                            v1 -= v1
                        default:
                            var v14 int16 = 42
                            _ = v14
                        }
                        _ = v13
                    }
                    v15 := GenS1{F0: 673}
                    _ = v15
                }
                v16 := v2[4294967295]
                v3 = v3
                _ = v16
            }
            v1 = v1
            _ = fi10
        }
    } else {
        var v17 int32 = -2147483648
        _ = v17
    }
    fmt.Println(v3, v3)
    v1 -= v1
    if false {
        v1 = v1
        var v18 int8 = 0
        _ = v18
    }
    v3 = 0
    v3 = v3
    v19 := v2[65535]
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
