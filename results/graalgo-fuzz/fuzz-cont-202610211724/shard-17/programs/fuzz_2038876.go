package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint64, p2 uint64) uint {
    v7 := []int{-443, 4294967295, 32768, 0}
    switch ("foo bar" + "hello") {
    case "hello":
        if (3.14159 != (1.5 * (-2.25 - 1.5))) {
            var v8 uint8 = 42
            _ = v8
        } else {
            var v9 float64 = (-2.25 / 0.1)
            fmt.Println(p2)
            _ = v9
        }
    case "":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case "a":
        var v10 bool = false
        _ = v10
    default:
        p1 += 42
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            if (241 >= 503) {
                var v13 int = fi12
                v14 := GenS1{F0: 2147483647, F1: 0}
                switch (fi11 << ((169) & 63)) {
                case 2147483647:
                    fmt.Println(fi12, fi12)
                default:
                    var v15 uint8 = 128
                    v15 = (^(v15))
                    _ = v15
                }
                _ = v13
                _ = v14
            }
            _ = fi12
        }
        v7 = append(v7, 42)
        _ = fi11
    }
    fmt.Println(p2, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 uint64, p1 bool, p2 uint32) int32 {
    fmt.Println(p2)
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            fmt.Println(fi16, fi16)
            _ = fi17
        }
        v18 := GenS2{F0: p2}
        _ = fi16
        _ = v18
    }
    var v19 uint = 2
    v20 := GenS3{F0: p0, F1: 180}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 2147483647
}

func main() {
    var v1 int8 = 127
    v2 := make(map[int]string)
    var v3 uint = 42
    var v4 int32 = 2147483647
    var v5 int8 = 127
    var v6 int8 = (v1 << ((2) & 63))
    var v21 int = 662
    var v22 int16 = 0
    var v23 uint32 = 42
    fmt.Println(v4, v4)
    v24 := GenS4{F0: 873}
    if false {
        if (0 > 1) {
            var v25 int16 = int16(v4)
            if (0 > v23) {
                var v26 int64 = 9223372036854775807
                switch (("a" + "\n\t") + ("" + "\n\t")) {
                case "a":
                    _ = len(v2)
                }
                v27 := GenS2{F0: (v23 >> ((v3) & 63))}
                _ = v26
                _ = v27
            }
            if ((v5 - v1) > 127) {
                for fi28 := 0; fi28 < 19; fi28++ {
                    v29 := GenS4{F0: 2}
                    _ = fi28
                    _ = v29
                }
                if (476 == 1) {
                    var v30 int32 = (v4 / (((^(v4))) | 1))
                    _ = v30
                }
            } else {
                fmt.Println(v21, v5, v4)
            }
            _ = v25
        }
        v1 = v1
        for fi31 := 0; fi31 < 12; fi31++ {
            for fi32 := 0; fi32 < 7; fi32++ {
                _ = len(v2)
                v33 := []int{}
                v5 = (v5 - -128)
                _ = fi32
                _ = v33
            }
            _ = fi31
        }
    } else {
        fmt.Println(v21)
        if (uint32(v3) < v23) {
            for fi34 := 0; fi34 < 19; fi34++ {
                v35 := GenS2{F0: (v23 >> (((v3 * 0)) & 63))}
                _ = fi34
                _ = v35
            }
            v1 = v1
        }
    }
    var v36 uint32 = (^(v23))
    v23 = v36
    var v37 uint64 = 18446744073709551615
    v38 := GenS1{F0: (v4 ^ v4), F1: 42}
    v5 = v1
    switch v21 {
    case 42:
        if true {
            var v39 int8 = v6
            var v40 uint = (v3 - v3)
            var v41 int64 = 441
            _ = v39
            _ = v40
            _ = v41
        }
    default:
        v42 := GenS4{F0: v37}
        _ = v42
    }
    var v43 string = "\n\t"
    var v44 uint8 = 161
    var v45 int64 = -1
    v43 = v43
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
