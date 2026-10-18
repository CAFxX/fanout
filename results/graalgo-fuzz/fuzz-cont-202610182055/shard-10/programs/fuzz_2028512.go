package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int8 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 int, p2 bool) uint16 {
    v6 := []int{0}
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 string = "hello"
            p2 = p2
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 128, F1: 255, F2: -77}
        if p2 {
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 uint = 2
                if p2 {
                    if false {
                        v6 = append(v6, p1)
                    }
                    p0 = 255
                }
                _ = fi11
                _ = v12
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            p1 = fi7
        } else {
            v6 = append(v6, 1)
            var v13 string = "a"
            _ = v13
        }
        _ = fi7
        _ = v10
    }
    var v14 int = (p1 * (p1 << ((42) & 63)))
    p0 -= 127
    if p2 {
        switch 128 {
        case 267:
            p2 = (0 <= 0)
        case 678:
            if p2 {
                if len(v6) > 0 {
                    v6[0] = v14
                }
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 uint8 = (p0 + p0)
                    v17 := GenS1{F0: p0, F1: 42, F2: -90}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v18 int32 = 727
                _ = v18
            }
        case 127:
            fmt.Println(p2)
        }
        fmt.Println(p1, p2)
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v6
    return 629
}
func genFn1(p0 uint8, p1 uint64) uint {
    var v19 uint16 = 255
    for fi20 := 0; fi20 < 9; fi20++ {
        if true {
            v19 = (v19 + (^(v19)))
            p0 -= 42
        }
        _ = fi20
    }
    p0 = (p0 * 157)
    _ = p0
    _ = p1
    _ = v19
    return 255
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    v3 := []string{("hello" + "hello"), (("\n\t" + "a") + ("a" + "a")), ("foo bar" + "a"), ((("foo bar" + "\n\t") + "hello") + "\n\t")}
    var v4 int8 = 42
    v5 := make(map[string]uint32)
    var v21 int16 = -32768
    if (65535 <= 364) {
        _ = len(v5)
        _ = len(v5)
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 bool = true
            v21 -= (v21 >> ((v2) & 63))
            v24 := GenS1{F0: 127, F1: (v2 / (((v2 & v2)) | 1)), F2: (^(v4))}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    fmt.Println(v21)
    if true {
        if (2147483647 != -2147483648) {
            _ = len(v1)
            v25 := GenS1{F0: 11, F1: genFn1(2, 0), F2: 60}
            _ = v25
        } else {
            v4 = -86
            v4 += (v4 * v4)
        }
    } else {
        var v26 uint8 = 50
        if false {
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 int8 = 127
                _ = len(v1)
                _ = fi27
                _ = v28
            }
        } else {
            _ = len(v3)
            var v29 int16 = 32767
            var v30 int32 = -1
            _ = v29
            _ = v30
        }
        _ = v26
    }
    if (2 != 65535) {
        _ = len(v3)
    }
    switch (("a" + ("foo bar" + "")) + "\n\t") {
    case "hello":
        var v31 float64 = 0.0
        _ = v31
    default:
        var v32 uint8 = 255
        _ = v32
    }
    if true {
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 int64 = 42
            fmt.Println(v34, v21)
            _ = fi33
            _ = v34
        }
    } else {
        v35 := GenS2{F0: 1}
        v4 = v4
        _ = v35
    }
    fmt.Println(v21, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v36 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
