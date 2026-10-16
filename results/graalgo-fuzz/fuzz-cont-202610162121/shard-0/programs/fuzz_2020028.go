package main

import (
    "fmt"
)

func genFn0(p0 bool, p1 uint16) uint {
    var v4 uint64 = 4294967295
    var v5 int16 = -1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 216
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    var v3 uint32 = 4294967295
    var v6 int64 = 0
    var v7 uint = 65535
    var v8 uint64 = 1
    fmt.Println(v8)
    if true {
        for fi9 := 0; fi9 < 8; fi9++ {
            fmt.Println(v8, fi9)
            _ = fi9
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            if true {
                var v11 int = -9223372036854775808
                _ = v11
            }
            _ = fi10
        }
    } else {
        var v12 bool = (false || ((v7 % (((v7 << (((v7 ^ genFn0(true, 65535))) & 63))) | 1)) >= genFn0(false, 255)))
        var v13 uint32 = 416
        if (true && v12) {
            v2 = 83
            var v14 int16 = 1000
            v3 += 545
            _ = v14
        }
        _ = v12
        _ = v13
    }
    switch ("a" + "a") {
    case "foo bar":
        if (1e10 < (3.14159 / 1.0)) {
            switch int(v6) {
            case 128:
                var v15 uint = uint(v3)
                _ = v15
            default:
                fmt.Println(v2)
            }
            for fi16 := 0; fi16 < 11; fi16++ {
                v7 -= genFn0(false, 42)
                v2 += 245
                if (true || (111 > v3)) {
                    var v17 uint = 18446744073709551615
                    v18 := []int{(fi16 & fi16), fi16, (^(fi16)), (fi16 - (fi16 - fi16))}
                    v7 += v7
                    _ = v17
                    _ = v18
                }
                _ = fi16
            }
        }
    case "":
        v6 += 1
    default:
        fmt.Println(v2, v8)
    }
    v6 = (v6 + (-(v6)))
    var v19 uint16 = 397
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
