package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int64 }
type GenS2 struct { F0 int; F1 string }

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    var v4 int = 9223372036854775807
    v4 = v4
    if true {
        v5 := GenS1{F0: (-(v4)), F1: 2147483647, F2: 1}
        for fi6 := 0; fi6 < 7; fi6++ {
            if true {
                var v7 uint16 = 32767
                if (-13 != -35) {
                    if true {
                        var v8 int64 = -9223372036854775808
                        _ = v8
                    }
                    var v9 int16 = -1
                    _ = v9
                }
                v10 := []int{v4, (v4 >> ((255) & 63)), fi6}
                _ = v10
                _ = v7
            } else {
                for fi11 := 0; fi11 < 12; fi11++ {
                    switch 4294967295 {
                    case -728:
                        if (false || false) {
                            var v12 string = "foo bar"
                            v4 += fi6
                            v4 -= fi6
                            _ = v12
                        }
                    }
                    _ = fi11
                }
                if false {
                    var v13 int64 = 1
                    var v14 int64 = v13
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 uint16 = 255
                    var v17 int = (fi15 - (v4 % (((fi6 ^ (fi6 * (^(fi15))))) | 1)))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            v18 := []int{}
            _ = fi6
            _ = v18
        }
        if false {
            if true {
                var v19 float64 = 0.1
                var v20 uint64 = (v1 << ((255) & 63))
                _ = v19
                _ = v20
            } else {
                var v21 uint = 255
                _ = v21
            }
            var v22 int = 42
            v2[(v4 & 9223372036854775807)] = 2
            _ = v22
        }
        _ = v5
    }
    _ = len(v3)
    v4 = 128
    _ = len(v3)
    var v23 uint32 = 0
    v24 := GenS1{F0: -439, F1: (v23 / (((v23 % (((v23 & v23)) | 1))) | 1)), F2: -919}
    v3[v23] = (("foo bar" + ("hello" + (("a" + "") + "a"))) + "\n\t")
    if false {
        var v25 bool = true
        _ = v25
    }
    switch v4 {
    case 156:
        for fi26 := 0; fi26 < 5; fi26++ {
            v27 := []int{}
            fmt.Println(v4, v23, v1)
            _ = fi26
            _ = v27
        }
    }
    if true {
        fmt.Println(v23)
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 int = 827
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: 128, F1: "a"}
    v4 -= v4
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
}
