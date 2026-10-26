package main

import (
    "fmt"
)

func genFn0() int64 {
    var v6 int8 = -87
    var v7 string = (("a" + "\n\t") + "foo bar")
    var v8 int8 = v6
    v8 -= -128
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    var v3 bool = true
    v4 := make(map[string]int)
    v5 := make(map[int]bool)
    v3 = false
    var v9 int = -436
    v3 = ("" != ("\n\t" + "foo bar"))
    v10 := []int{v9}
    v3 = v3
    if false {
        switch 2147483647 {
        case 32768:
            v3 = false
        case -589:
            if v3 {
                fmt.Println(v9)
                var v11 float32 = float32((1.5 + 1.5))
                _ = v11
            } else {
                fmt.Println(v3)
                for fi12 := 0; fi12 < 4; fi12++ {
                    v3 = (99 < -1)
                    var v13 int16 = 32767
                    var v14 uint = 648
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        case 751:
            delete(v1, 4294967295)
        default:
            if true {
                for fi15 := 0; fi15 < 7; fi15++ {
                    var v16 int32 = -2147483648
                    var v17 string = ""
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                _ = len(v10)
                for fi18 := 0; fi18 < 9; fi18++ {
                    if true {
                        v3 = (255 != 255)
                    }
                    v9 -= fi18
                    if false {
                        v3 = (315 <= 1)
                        v2[42] = 0
                    } else {
                        v3 = v3
                    }
                    _ = fi18
                }
            }
        }
    } else {
        var v19 uint64 = 4294967295
        _ = v19
    }
    v3 = (42 > 32767)
    var v20 bool = v3
    _ = len(v10)
    var v21 bool = true
    for fi22 := 0; fi22 < 13; fi22++ {
        v3 = (v20 && v21)
        var v23 int64 = genFn0()
        _ = len(v10)
        _ = fi22
        _ = v23
    }
    fmt.Println(v9, v9, v3)
    var v24 int8 = -128
    var v25 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
