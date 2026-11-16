package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0() int32 {
    var v7 int32 = -2147483648
    v7 -= (v7 << ((794) & 63))
    if true {
        v7 = (v7 ^ v7)
        var v8 uint64 = 2
        _ = v8
    } else {
        v9 := GenS1{F0: -108, F1: 15}
        if true {
            fmt.Println(v7)
            var v10 uint64 = 18446744073709551615
            v10 -= 2
            _ = v10
        } else {
            v11 := GenS1{F0: (v7 ^ v7), F1: 114}
            var v12 uint8 = 128
            _ = v11
            _ = v12
        }
        var v13 int64 = 1
        _ = v13
        _ = v9
    }
    var v14 string = ("hello" + "a")
    fmt.Println(v14, v7)
    _ = v14
    _ = v7
    return 632
}
func genFn1(p0 int64, p1 int, p2 uint16) int64 {
    v15 := []int{(p1 + (p1 + 34)), 1, (p1 | p1)}
    if (65535 < 255) {
        switch "" {
        case "a":
            for fi16 := 0; fi16 < 17; fi16++ {
                if ("\n\t" >= "a") {
                    v15 = append(v15, fi16)
                }
                for fi17 := 0; fi17 < 1; fi17++ {
                    if true {
                        v18 := GenS1{F0: 2147483647, F1: 222}
                        var v19 uint8 = 190
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
                _ = fi16
            }
        case "foo bar":
            var v20 int64 = 9223372036854775807
            _ = v20
        case "\n\t":
            var v21 int = 1
            _ = v21
        }
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (^(p0))
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 255
    v3 := make(map[uint32]string)
    v4 := []int{}
    var v5 bool = (0.1 > 3.14159)
    var v6 uint64 = 1
    for fi22 := 0; fi22 < 5; fi22++ {
        if v5 {
            var v23 int8 = 48
            fmt.Println(v2)
            _ = v23
        } else {
            v24 := GenS1{F0: -1, F1: v2}
            for fi25 := 0; fi25 < 4; fi25++ {
                if len(v4) > 0 {
                    v4[0] = (fi25 * (fi25 & (^(fi22))))
                }
                v26 := v3[732]
                v2 += (^(v2))
                _ = fi25
                _ = v26
            }
            if (v5 && (2147483647 == -2147483648)) {
                v5 = true
                _ = len(v1)
                var v27 int = (^(fi22))
                _ = v27
            }
            _ = v24
        }
        _ = fi22
    }
    if v5 {
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 uint8 = (v2 >> ((344) & 63))
            var v30 int64 = -9223372036854775808
            switch (fi28 >> ((255) & 63)) {
            case 32768:
                for fi31 := 0; fi31 < 12; fi31++ {
                    fmt.Println(fi28, v2, v2)
                    v30 += v30
                    _ = fi31
                }
            case 1:
                v29 -= v2
            case 2147483647:
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 int16 = 0
                    v30 -= -848
                    var v34 uint = 657
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            _ = fi28
            _ = v29
            _ = v30
        }
    } else {
        v35 := GenS1{F0: 2147483647, F1: (v2 / ((v2) | 1))}
        _ = v35
    }
    fmt.Println(v2, v5)
    var v36 int64 = -1
    var v37 float64 = 0.0
    var v38 uint8 = 42
    for fi39 := 0; fi39 < 17; fi39++ {
        if false {
            v40 := GenS1{F0: 42, F1: v38}
            _ = v40
        }
        _ = fi39
    }
    var v41 float32 = 1e10
    v38 += 1
    v6 = (v6 | v6)
    v36 = (v36 ^ v36)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v42 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
