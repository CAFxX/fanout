package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint64 }
func genFn0(p0 uint64, p1 int) uint {
    switch ((("hello" + "") + ("" + "hello")) + "foo bar") {
    case "":
        p0 += (p0 * (p0 % ((p0) | 1)))
    }
    v7 := []int{(p1 | 568), (p1 / ((p1) | 1)), p1, (p1 - (-(p1)))}
    var v8 int16 = 134
    var v9 uint16 = 32767
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 1
    var v3 int = 32767
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    var v6 string = "a"
    v10 := []int{0}
    v1 = (v1 - genFn0(119, v3))
    for fi11 := 0; fi11 < 2; fi11++ {
        v3 = (fi11 | (-(fi11)))
        for fi12 := 0; fi12 < 4; fi12++ {
            v1 = v1
            v3 -= (v3 << (((^(v1))) & 63))
            if false {
                fmt.Println(v2, v3, fi11)
                v13 := GenS1{F0: (v3 / (((v3 | fi12)) | 1)), F1: true, F2: 4294967295}
                _ = v13
            }
            if (2 > v1) {
                var v14 int = (^(v3))
                switch (fi11 ^ v14) {
                case 2147483647:
                    var v15 uint32 = 0
                    _ = v15
                case 194:
                    fmt.Println(v14, v1)
                }
                _ = v14
            }
            _ = fi12
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int = 845
            _ = fi16
            _ = v17
        }
        _ = fi11
    }
    v18 := GenS1{F0: 127, F1: ((-520 <= -693) && true), F2: 1}
    v6 = v6
    v2 -= v2
    var v19 uint = (v1 >> ((genFn0(4294967295, v3)) & 63))
    _ = len(v5)
    fmt.Println(v19, v19, v2)
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v3, v3)
        fmt.Println(v6, v3, v1)
        _ = fi20
    }
    v4[int(v2)] = ("hello" + v6)
    v2 = -705
    _ = len(v10)
    v3 = (v3 << ((358) & 63))
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 uint = genFn0(99, 593)
        fmt.Println(v2, v1)
        if false {
            var v23 uint8 = 99
            if false {
                v5[2147483648] = v6
                _ = len(v10)
            }
            _ = v23
        } else {
            if (false || (1 > 2147483647)) {
                fmt.Println(v3, v2, fi21)
                v22 = v1
                fmt.Println(v19)
            }
            var v24 int64 = -1
            var v25 string = "foo bar"
            _ = v24
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    delete(v4, (-(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
