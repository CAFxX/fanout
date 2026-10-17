package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 string; F2 uint8 }

func main() {
    var v1 uint = 255
    var v2 uint16 = 32767
    v3 := []string{("" + "hello"), (("a" + ("a" + "foo bar")) + "\n\t")}
    var v4 bool = (2147483647 > 421)
    v5 := make(map[int]string)
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint16 = v2
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v1 += v1
        var v10 int16 = 32767
        _ = fi9
        _ = v10
    }
    switch "" {
    case "foo bar":
        for fi11 := 0; fi11 < 10; fi11++ {
            if ((!(!(!(!true)))) || v4) {
                v12 := GenS1{F0: 255}
                v4 = true
                fmt.Println(v2)
                _ = v12
            }
            _ = fi11
        }
    case "hello":
        var v13 uint32 = uint32(v1)
        _ = v13
    }
    if ((42 > -1) || (127 > 42)) {
        fmt.Println(v2)
        var v14 uint64 = 900
        _ = v14
    } else {
        v1 = 232
        fmt.Println(v1, v4)
    }
    v15 := GenS2{F0: 2147483648}
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS3{F0: 0}
        _ = fi16
        _ = v17
    }
    if v4 {
        if (!(!false)) {
            v4 = v4
        } else {
            if false {
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 float64 = (3.14159 / 1e10)
                    var v20 int8 = 42
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v1, v4)
            }
        }
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int32 = -1
            _ = fi21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        v1 -= (v1 & (v1 >> ((2) & 63)))
        if len(v3) > 0 {
            v3[0] = ("hello" + "a")
        }
        if true {
            v3 = append(v3, "a")
        } else {
            v1 = v1
            v1 = 2
            if v4 {
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS4{F0: 762}
                    fmt.Println(v4, v4, fi24)
                    v4 = (true || (1 <= -691))
                    _ = fi24
                    _ = v25
                }
                v26 := GenS5{F0: 128}
                v27 := GenS6{F0: 2, F1: ("hello" + "foo bar"), F2: 9}
                _ = v26
                _ = v27
            }
        }
        _ = fi23
    }
    v2 += 255
    v2 += (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
