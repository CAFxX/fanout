package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8) int {
    v4 := GenS1{F0: 2147483647}
    for fi5 := 0; fi5 < 9; fi5++ {
        p0 = p0
        fmt.Println(fi5, p0)
        if true {
            switch -404 {
            case 32767:
                for fi6 := 0; fi6 < 2; fi6++ {
                    var v7 int8 = (p0 * (p0 & p0))
                    var v8 uint64 = 4294967295
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            case 687:
                var v9 uint8 = 0
                _ = v9
            default:
                var v10 string = "foo bar"
                _ = v10
            }
            switch "" {
            case "foo bar":
                v11 := []int{-98, -2147483648}
                _ = v11
            }
            if true {
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 bool = ((false && (!false)) && false)
                    p0 += p0
                    var v14 uint32 = 655
                    var v15 int = (fi12 >> ((701) & 63))
                    if false {
                        var v16 bool = false
                        _ = v16
                    } else {
                        fmt.Println(fi5, fi5, v15)
                        v13 = (28 != 2)
                    }
                    v13 = ((!v13) || true)
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                v17 := GenS2{F0: 183}
                p0 = (p0 & (p0 << ((2) & 63)))
                _ = v17
            } else {
                fmt.Println(fi5)
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 uint = 255
                    if (1 != p0) {
                        var v20 int64 = -1
                        fmt.Println(v19, fi18, fi5)
                        var v21 int16 = -606
                        _ = v20
                        _ = v21
                    }
                    var v22 uint64 = 4294967295
                    var v23 int = int(v19)
                    _ = fi18
                    _ = v19
                    _ = v22
                    _ = v23
                }
            }
        }
        _ = fi5
    }
    p0 += (p0 << ((0) & 63))
    _ = p0
    _ = v4
    return 686
}

func main() {
    var v1 uint = 2
    v2 := []string{("\n\t" + ""), ("foo bar" + "foo bar"), ""}
    v3 := make(map[uint32]uint32)
    fmt.Println(v1)
    var v24 uint16 = 32767
    v1 += v1
    v24 -= v24
    fmt.Println(v1)
    delete(v3, 391)
    var v25 int8 = 42
    var v26 uint = (^(v1))
    delete(v3, uint32(v24))
    v27 := GenS2{F0: 18}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
