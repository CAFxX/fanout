package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32; F2 string }
type GenS3 struct { F0 int16; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 42
    v3 := make(map[string]uint32)
    var v4 int = -181
    var v5 uint8 = 2
    var v6 int8 = 127
    fmt.Println(v4)
    v6 = -1
    for fi7 := 0; fi7 < 20; fi7++ {
        v4 += fi7
        _ = fi7
    }
    v8 := GenS1{F0: 42}
    switch "\n\t" {
    case "":
        switch 747 {
        case 591:
            v2 = (^(v2))
        case 172:
            v1["hello"] = "\n\t"
        default:
            if ((0.0 <= float64((0.0 + (1e10 + 0.1)))) || false) {
                v5 = 1
            } else {
                switch (^(v4)) {
                case -1:
                    v4 -= v4
                default:
                    v3["a"] = 4294967295
                }
            }
        }
    case "foo bar":
        fmt.Println(v2, v2)
    case "hello":
        var v9 float64 = 0.1
        _ = v9
    default:
        fmt.Println(v5, v2, v2)
    }
    var v10 bool = (!true)
    if (v5 >= v5) {
        v2 += 42
        var v11 uint16 = 1
        if v10 {
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 uint64 = v2
                if v10 {
                    v5 += (v5 & (v5 / ((v5) | 1)))
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 uint16 = 0
                    _ = fi14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            delete(v1, "")
        }
        _ = v11
    } else {
        v4 += (-(v4))
    }
    var v16 int = -92
    var v17 uint = 1
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 6; fi19++ {
            v17 -= 255
            var v20 float32 = (3.14159 - ((0.0 * -2.25) + 3.14159))
            v5 += v5
            _ = fi19
            _ = v20
        }
        switch fi18 {
        case 1:
            var v21 bool = (1 < 42)
            _ = v21
        case 994:
            v22 := []int{int(v6), (v4 % ((v4) | 1)), fi18}
            _ = v22
        case 32768:
            if (v2 >= v2) {
                v6 += 42
                v5 = (v5 + (v5 & (^(v5))))
                switch ("hello" + "") {
                case "foo bar":
                    v23 := []int{610, v16, v16}
                    _ = v23
                case "a":
                    v24 := GenS1{F0: 2}
                    _ = v24
                case "\n\t":
                    var v25 int = int(v6)
                    _ = v25
                }
            }
        default:
            var v26 bool = (v5 != v5)
            _ = v26
        }
        var v27 string = "hello"
        _ = fi18
        _ = v27
    }
    switch (v16 + v16) {
    case -2147483648:
        switch "hello" {
        case "foo bar":
            for fi28 := 0; fi28 < 12; fi28++ {
                v4 = (v16 & (fi28 + (v16 | v4)))
                _ = fi28
            }
        case "a":
            v6 = v6
        default:
            var v29 uint16 = 210
            _ = v29
        }
    case -356:
        if v10 {
            var v30 uint8 = v5
            v31 := GenS2{F0: 1000, F1: 42, F2: "foo bar"}
            _ = v30
            _ = v31
        }
    case 127:
        if ((uint(v2) == 0) || v10) {
            v32 := GenS1{F0: 70}
            _ = v32
        }
    default:
        v4 = v4
    }
    v33 := GenS1{F0: 2147483647}
    if v10 {
        for fi34 := 0; fi34 < 18; fi34++ {
            v1["a"] = (("\n\t" + "") + "a")
            v5 -= (v5 << ((v17) & 63))
            _ = fi34
        }
        v6 = (^(v6))
    }
    v16 = v4
    for fi35 := 0; fi35 < 17; fi35++ {
        if (v10 || true) {
            var v36 int8 = 100
            _ = v36
        } else {
            var v37 bool = true
            v38 := GenS1{F0: uint32(v2)}
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 13; fi39++ {
            v40 := []int{v16, (v4 % ((v16) | 1)), v4}
            v10 = true
            _ = fi39
            _ = v40
        }
        _ = fi35
    }
    v6 -= v6
    for fi41 := 0; fi41 < 15; fi41++ {
        if false {
            v42 := GenS3{F0: 1000, F1: (^(v17)), F2: 2147483648}
            var v43 int64 = -1
            _ = v42
            _ = v43
        }
        v16 = (v16 | v4)
        _ = fi41
    }
    fmt.Println(v2, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
