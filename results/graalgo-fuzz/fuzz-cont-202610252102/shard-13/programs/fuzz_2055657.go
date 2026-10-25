package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint16 {
    switch 1 {
    case 32768:
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 float32 = (float32(-2.25) / 1e10)
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS1{F0: false}
                _ = fi8
                _ = v9
            }
            var v10 uint64 = 2
            _ = fi6
            _ = v10
            _ = v7
        }
    case 4294967295:
        var v11 uint = 65535
        _ = v11
    case 1:
        v12 := GenS1{F0: (!false)}
        _ = v12
    default:
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 uint16 = 65535
            _ = fi13
            _ = v14
        }
    }
    var v15 bool = (false || ((559 <= 311) || ((!(440 == 9223372036854775807)) && true)))
    _ = v15
    return 31
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    v3 := []int{}
    v4 := make(map[uint32]bool)
    var v5 int64 = 601
    var v16 int64 = (v5 / ((42) | 1))
    v5 = v16
    delete(v4, 2147483648)
    switch ("\n\t" + "foo bar") {
    case "hello":
        v17 := GenS1{F0: true}
        _ = v17
    case "a":
        var v18 bool = (true && (1e10 <= 0.1))
        _ = v18
    case "foo bar":
        v5 -= -1
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if ("a" != "a") {
            for fi20 := 0; fi20 < 20; fi20++ {
                v16 -= (v5 ^ (v5 << ((637) & 63)))
                var v21 int32 = -2147483648
                v22 := GenS1{F0: (true && false)}
                fmt.Println(fi20, v5)
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := GenS1{F0: v2}
            switch "foo bar" {
            case "a":
                var v24 int8 = 19
                _ = v24
            case "foo bar":
                _ = len(v1)
            case "":
                v25 := v1[0]
                _ = v25
            default:
                var v26 float32 = (1.5 / (0.0 - (float32((-2.25 * (3.14159 - (((0.1 * 0.0) / (((3.14159 / 1.5) * (0.1 + 0.1)) - 0.0)) / (-2.25 * 3.14159))))) / float32(0.1))))
                _ = v26
            }
            _ = v23
        }
        _ = fi19
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        v16 = 943
        _ = fi27
    }
    v16 = 9223372036854775807
    v28 := GenS1{F0: (false && v2)}
    v3 = append(v3, 0)
    switch "\n\t" {
    case "hello":
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 int32 = 868
            switch fi29 {
            case 2147483647:
                var v31 uint8 = 255
                _ = v31
            default:
                var v32 uint32 = 2
                _ = v32
            }
            var v33 uint16 = 32767
            _ = fi29
            _ = v30
            _ = v33
        }
    case "foo bar":
        fmt.Println(v16)
    default:
        for fi34 := 0; fi34 < 17; fi34++ {
            var v35 uint8 = 1
            _ = fi34
            _ = v35
        }
    }
    fmt.Println(v16, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
