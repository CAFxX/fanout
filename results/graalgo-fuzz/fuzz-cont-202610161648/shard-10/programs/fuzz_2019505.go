package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint32, p2 uint) bool {
    v6 := []int{-850, 4294967295, 4294967295, -250}
    p1 = 978
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}

func main() {
    var v1 bool = false
    var v2 int = 127
    v3 := []uint32{492, 4294967295, 4294967295}
    v4 := []int{(v2 ^ (v2 * v2))}
    v5 := make(map[string]string)
    v1 = v1
    v7 := GenS1{F0: v2}
    switch "hello" {
    case "foo bar":
        _ = len(v5)
    }
    v8 := GenS2{F0: 656, F1: 2}
    v1 = ((((!(244 >= 0)) && v1) && v1) && (false && (false && (255 > 2))))
    if v1 {
        v1 = v1
        if v1 {
            var v9 int16 = 0
            _ = v9
        } else {
            fmt.Println(v1)
            var v10 uint64 = 578
            _ = v10
        }
        v11 := GenS3{F0: 18446744073709551615}
        _ = v11
    } else {
        switch (v2 << ((1) & 63)) {
        case 213:
            var v12 uint8 = 255
            _ = v12
        case 32768:
            if (18446744073709551615 == 4294967295) {
                var v13 bool = (("" + "a") <= (("hello" + ((("a" + "\n\t") + "a") + "\n\t")) + "a"))
                var v14 int16 = 0
                v13 = v13
                _ = v13
                _ = v14
            } else {
                if v1 {
                    var v15 string = "hello"
                    _ = v15
                } else {
                    var v16 uint = 2
                    fmt.Println(v16)
                    fmt.Println(v1, v1)
                    _ = v16
                }
                if v1 {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            if (2 > 712) {
                                v3 = append(v3, 4294967295)
                                v3 = append(v3, 1)
                                v19 := v5["a"]
                                _ = v19
                            }
                            v1 = v1
                            var v20 int = fi17
                            _ = fi18
                            _ = v20
                        }
                        v1 = (42 != -1)
                        _ = fi17
                    }
                    var v21 int8 = 55
                    var v22 int64 = 598
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 4; fi23++ {
                    _ = len(v4)
                    _ = fi23
                }
            }
        }
        v24 := GenS3{F0: 369}
        _ = v24
    }
    var v25 uint8 = 42
    var v26 float32 = (0.1 * (1.5 + 1e10))
    var v27 uint16 = 78
    v2 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
