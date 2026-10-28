package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
func genFn0(p0 bool, p1 int32) int {
    p0 = false
    var v5 uint32 = 197
    if (156 >= 383) {
        p1 += p1
        v6 := []int{2147483647}
        _ = v6
    } else {
        v7 := []int{4294967295, 4294967295, -401}
        v7 = append(v7, 127)
        _ = v7
    }
    v8 := []int{-1, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 205
}
func genFn1(p0 uint8) uint16 {
    fmt.Println(p0)
    if true {
        if false {
            if true {
                p0 = 255
            } else {
                var v9 int16 = -1
                if (-1 > 128) {
                    v10 := GenS1{F0: 1}
                    v9 += v9
                    _ = v10
                } else {
                    v11 := GenS2{F0: 0}
                    p0 = p0
                    _ = v11
                }
                if false {
                    p0 -= 255
                }
                _ = v9
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                for fi13 := 0; fi13 < 20; fi13++ {
                    fmt.Println(p0)
                    _ = fi13
                }
                _ = fi12
            }
        } else {
            fmt.Println(p0)
            v14 := []int{127, -612, -690}
            _ = v14
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 int16 = -979
            _ = fi15
            _ = v16
        }
    } else {
        v17 := []int{}
        var v18 uint = 65535
        _ = v17
        _ = v18
    }
    var v19 uint8 = p0
    switch "" {
    case "a":
        p0 = (^(p0))
    }
    _ = p0
    _ = v19
    return 1
}

func main() {
    var v1 bool = false
    var v2 uint16 = 645
    v3 := []uint32{2, 2, 2}
    var v4 float32 = 1.5
    var v20 bool = v1
    if v1 {
        var v21 int = 32768
        _ = v21
    }
    _ = len(v3)
    var v22 int16 = -32768
    var v23 uint32 = 4294967295
    if ("foo bar" > "") {
        fmt.Println(v20, v1, v4)
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := GenS1{F0: (^(v23))}
        _ = fi24
        _ = v25
    }
    if v1 {
        v23 += (v23 ^ (v23 ^ v23))
        v3 = append(v3, 497)
    }
    switch "foo bar" {
    case "a":
        if (282 != 1) {
            for fi26 := 0; fi26 < 1; fi26++ {
                switch ("\n\t" + (("hello" + "hello") + ("foo bar" + "\n\t"))) {
                case "foo bar":
                    var v27 uint = 1
                    _ = v27
                default:
                    var v28 int64 = -9223372036854775808
                    _ = v28
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v22 = 0
                _ = fi26
            }
            v20 = (genFn1(0) >= v2)
            var v29 int32 = 42
            _ = v29
        } else {
            for fi30 := 0; fi30 < 11; fi30++ {
                switch fi30 {
                case 42:
                    v3 = append(v3, 2147483648)
                default:
                    _ = len(v3)
                }
                v3 = append(v3, (v23 & (v23 >> ((255) & 63))))
                fmt.Println(fi30, v20, fi30)
                _ = fi30
            }
            switch 2147483647 {
            case 9223372036854775807:
                var v31 uint16 = (^(v2))
                _ = v31
            }
            for fi32 := 0; fi32 < 9; fi32++ {
                v3 = append(v3, v23)
                var v33 int8 = 127
                fmt.Println(v20, fi32)
                _ = fi32
                _ = v33
            }
        }
    }
    var v34 bool = false
    for fi35 := 0; fi35 < 16; fi35++ {
        v22 -= -176
        _ = fi35
    }
    var v36 uint64 = 18446744073709551615
    var v37 uint = 2
    if ((!v20) && (255 != 1000)) {
        v3 = append(v3, v23)
        fmt.Println(v34, v37)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v4
}
