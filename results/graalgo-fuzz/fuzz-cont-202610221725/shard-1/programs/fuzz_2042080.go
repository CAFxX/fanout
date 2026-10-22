package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint64) int64 {
    v5 := []int{}
    p1 = p1
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint16 = 32767
        p1 -= (^(p1))
        _ = fi6
        _ = v7
    }
    if (42 > 32767) {
        if ((p1 - (p1 * (^(p1)))) == p1) {
            var v8 uint64 = 1
            _ = v8
        }
        var v9 uint = 449
        fmt.Println(p1)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            if len(v5) > 0 {
                v5[0] = fi10
            }
            _ = fi10
        }
        var v11 int64 = -9223372036854775808
        _ = v11
    }
    var v12 uint32 = 545
    _ = p0
    _ = p1
    _ = v12
    _ = v5
    return int64(p1)
}

func main() {
    v1 := []string{("" + "\n\t"), ("a" + "hello")}
    var v2 float32 = 1e10
    var v3 int32 = 567
    v4 := make(map[uint32]bool)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint32 = 2
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int = fi13
            _ = fi15
            _ = v16
        }
        v14 += v14
        _ = fi13
        _ = v14
    }
    _ = len(v4)
    v17 := GenS1{F0: 585}
    for fi18 := 0; fi18 < 17; fi18++ {
        v2 -= v2
        fmt.Println(v3)
        _ = fi18
    }
    fmt.Println(v2)
    var v19 string = "hello"
    var v20 uint8 = 66
    delete(v4, 65535)
    v21 := GenS1{F0: 4294967295}
    v22 := GenS1{F0: 42}
    var v23 uint32 = 4294967295
    v24 := GenS1{F0: (^(v23))}
    switch 52 {
    case 4294967295:
        switch -9223372036854775808 {
        case 992:
            var v25 float64 = (3.14159 - 1e10)
            _ = v25
        }
    case 128:
        switch ("" + v19) {
        case "hello":
            if false {
                v19 = ((v19 + v19) + "\n\t")
            } else {
                switch 32767 {
                case -1:
                    v1 = append(v1, (v19 + "\n\t"))
                }
            }
        case "foo bar":
            for fi26 := 0; fi26 < 6; fi26++ {
                for fi27 := 0; fi27 < 12; fi27++ {
                    var v28 uint8 = (v20 >> ((42) & 63))
                    switch "hello" {
                    case "":
                        _ = len(v1)
                    default:
                        v20 = 1
                    }
                    _ = len(v1)
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 1; fi29++ {
                    _ = len(v4)
                    var v30 string = ("hello" + v19)
                    v19 = v30
                    _ = fi29
                    _ = v30
                }
                _ = fi26
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
