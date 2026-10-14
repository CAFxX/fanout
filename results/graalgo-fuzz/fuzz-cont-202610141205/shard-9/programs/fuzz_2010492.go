package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 string, p1 int32, p2 uint) string {
    if true {
        p1 = p1
    } else {
        v4 := []int{-252, 704, 1, -612}
        var v5 int = 128
        _ = v4
        _ = v5
    }
    p2 = 390
    p2 += p2
    var v6 uint8 = 151
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "hello"
}
func genFn1(p0 uint16, p1 int8, p2 string) int32 {
    p2 = (("hello" + (p2 + p2)) + ((p2 + (("\n\t" + "\n\t") + (p2 + p2))) + (p2 + p2)))
    p1 -= 53
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]bool)
    var v3 int8 = (v1 % ((v1) | 1))
    fmt.Println(v3, v1)
    if false {
        var v7 string = "a"
        _ = v7
    } else {
        var v8 uint = 846
        v9 := []int{850, 2147483647, 361, 2147483647}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42}
    if false {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: 2}
            var v13 int64 = 464
            _ = fi11
            _ = v12
            _ = v13
        }
        v3 = (v1 ^ v3)
        var v14 float32 = 1e10
        _ = v14
    }
    var v15 int64 = 9223372036854775807
    fmt.Println(v3)
    var v16 int32 = genFn1(520, (v3 | -128), "a")
    if true {
        for fi17 := 0; fi17 < 7; fi17++ {
            switch (fi17 + fi17) {
            case 287:
                fmt.Println(v1)
            case 128:
                var v18 int64 = 89
                _ = v18
            }
            if (!false) {
                switch ("" + "") {
                case "":
                    v16 -= (v16 % ((v16) | 1))
                }
                var v19 int8 = (v3 / ((42) | 1))
                var v20 uint16 = 265
                _ = v19
                _ = v20
            }
            var v21 uint = 0
            _ = fi17
            _ = v21
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            for fi23 := 0; fi23 < 4; fi23++ {
                if false {
                    v1 = (v1 << ((613) & 63))
                    v16 = v16
                }
                var v24 int8 = (v3 ^ v3)
                _ = fi23
                _ = v24
            }
            v1 = v3
            var v25 int8 = 0
            _ = fi22
            _ = v25
        }
        for fi26 := 0; fi26 < 7; fi26++ {
            delete(v2, fi26)
            var v27 uint8 = 128
            fmt.Println(fi26, v3)
            _ = fi26
            _ = v27
        }
    } else {
        if true {
            for fi28 := 0; fi28 < 8; fi28++ {
                v29 := []int{2147483647, (fi28 * 42)}
                v3 -= 42
                v3 = v3
                _ = fi28
                _ = v29
            }
            var v30 int16 = 302
            _ = v30
        } else {
            if (0 >= 895) {
                switch 233 {
                case -349:
                    v31 := GenS1{F0: 597}
                    _ = v31
                default:
                    v32 := GenS1{F0: 1}
                    _ = v32
                }
            }
            if (false || true) {
                v33 := []int{}
                v34 := GenS1{F0: 0}
                _ = v33
                _ = v34
            } else {
                v15 = (v15 % ((v15) | 1))
                v15 = v15
                v1 += v1
            }
            if false {
                v35 := []int{}
                var v36 int64 = (v15 - (-(v15)))
                v37 := GenS1{F0: 2}
                _ = v35
                _ = v36
                _ = v37
            } else {
                switch "" {
                case "\n\t":
                    v38 := []int{137, 2147483647}
                    _ = v38
                default:
                    delete(v2, -2147483648)
                }
            }
        }
    }
    var v39 int32 = (v16 | (v16 >> ((235) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v39
}
