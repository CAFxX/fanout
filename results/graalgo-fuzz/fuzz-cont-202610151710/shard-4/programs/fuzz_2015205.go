package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int; F1 int32; F2 uint16 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]uint32)
    var v3 float64 = (3.14159 / 1.0)
    v4 := GenS1{F0: 509}
    v1 -= 154
    for fi5 := 0; fi5 < 4; fi5++ {
        v2["\n\t"] = 2
        _ = fi5
    }
    var v6 uint64 = uint64(v1)
    v7 := v2["foo bar"]
    switch 371 {
    case 4294967295:
        switch "a" {
        case "hello":
            if (0 != 0) {
                v8 := []int{2147483647, 32768, 53, 127}
                fmt.Println(v7, v3, v6)
                _ = v8
            } else {
                var v9 int32 = -1
                fmt.Println(v9, v9)
                var v10 string = ("" + (("hello" + "") + ("hello" + "a")))
                _ = v10
                _ = v9
            }
        default:
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := []int{fi11, (-(fi11)), (fi11 - fi11)}
                _ = fi11
                _ = v12
            }
        }
    default:
        v13 := GenS1{F0: v6}
        _ = v13
    }
    var v14 uint16 = 0
    fmt.Println(v1)
    for fi15 := 0; fi15 < 17; fi15++ {
        if (-791 < 947) {
            fmt.Println(fi15)
            var v16 float64 = v3
            for fi17 := 0; fi17 < 17; fi17++ {
                switch (("a" + "a") + "\n\t") {
                case "hello":
                    v1 += 128
                case "foo bar":
                    v18 := GenS1{F0: 0}
                    _ = v18
                }
                v19 := []int{}
                _ = fi17
                _ = v19
            }
            _ = v16
        } else {
            var v20 uint64 = (v6 & uint64(v14))
            var v21 int8 = 100
            _ = v20
            _ = v21
        }
        delete(v2, ("foo bar" + ""))
        _ = fi15
    }
    var v22 uint16 = v14
    v22 -= (v14 / ((v22) | 1))
    v1 = uint8(v14)
    v23 := GenS2{F0: (!true), F1: 2147483647}
    v24 := GenS3{F0: -1, F1: 42, F2: (v14 * (v22 + 32767))}
    fmt.Println(v1)
    v6 = (v6 & uint64(v14))
    switch 727 {
    case 2147483647:
        v6 = (^(v6))
    case 1:
        v25 := GenS3{F0: 4294967295, F1: 42, F2: v14}
        _ = v25
    case -2147483648:
        var v26 string = ((("\n\t" + "") + "hello") + "a")
        _ = v26
    default:
        if false {
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 uint16 = (v22 << ((0) & 63))
                _ = fi27
                _ = v28
            }
            v7 = (v7 * (v7 << ((uint(v1)) & 63)))
            var v29 int8 = 35
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
