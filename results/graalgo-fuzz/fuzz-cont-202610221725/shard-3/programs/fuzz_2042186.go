package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int32 = -859
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 200
    v8 := []int{32767, -1, 184, -977}
    switch 317 {
    case 128:
        var v9 bool = (true && true)
        _ = v9
    }
    v7 = (v7 & 1)
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint8 = 236
    var v2 bool = (((!(128 < v1)) || true) || false)
    var v3 int = 128
    var v4 uint8 = (v1 * 163)
    switch (v3 >> ((18446744073709551615) & 63)) {
    case 4294967295:
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS1{F0: "foo bar"}
            v1 = 212
            var v12 float32 = (0.0 * 1e10)
            _ = fi10
            _ = v11
            _ = v12
        }
    case -307:
        fmt.Println(v3, v1)
    }
    var v13 uint64 = 4294967295
    var v14 uint64 = (v13 & 18446744073709551615)
    if ((2 == v14) && (-9223372036854775808 != 1)) {
        var v15 float64 = (((((float64(((0.0 - ((0.0 - 1.5) * (float32(1.5) + (float32(float64(-2.25)) / 0.1)))) - 1e10)) * float64(1.5)) / float64((-2.25 - float32((0.1 / 0.1))))) / 1.5) + 0.1) / ((1.5 + float64(0.1)) / float64(0.1)))
        fmt.Println(v1, v1)
        _ = v15
    } else {
        switch "" {
        case "\n\t":
            var v16 uint = 835
            _ = v16
        }
    }
    v1 -= 143
    v4 = (v4 & (v1 | 36))
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            fmt.Println(v14, fi18, v3)
            _ = fi18
        }
        v19 := GenS1{F0: "\n\t"}
        _ = fi17
        _ = v19
    }
    if true {
        v20 := []int{v3}
        v1 -= uint8(v14)
        _ = v20
    }
    switch "\n\t" {
    case "":
        v21 := []int{128, 32768, (v3 / (((^(v3))) | 1)), 128}
        _ = v21
    case "\n\t":
        fmt.Println(v13)
    case "foo bar":
        v22 := GenS2{F0: 268}
        _ = v22
    default:
        var v23 int64 = -9223372036854775808
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        if (3.14159 == 3.14159) {
            for fi25 := 0; fi25 < 9; fi25++ {
                v1 = (v4 >> ((0) & 63))
                var v26 uint = 65535
                _ = fi25
                _ = v26
            }
        } else {
            v2 = ("hello" != "")
            var v27 string = "\n\t"
            fmt.Println(v14, v4, v13)
            _ = v27
        }
        if v2 {
            v4 -= v4
            for fi28 := 0; fi28 < 16; fi28++ {
                if (-890 != -1) {
                    var v29 int16 = 1
                    _ = v29
                } else {
                    var v30 uint32 = 633
                    var v31 float32 = (1.5 + 3.14159)
                    v32 := []int{fi28, 26, (v3 & fi28), (fi24 | int(v13))}
                    _ = v30
                    _ = v31
                    _ = v32
                }
                v2 = true
                _ = fi28
            }
            var v33 int64 = 0
            _ = v33
        }
        v13 += v14
        _ = fi24
    }
    var v34 uint64 = (v14 / ((v14) | 1))
    v4 += 42
    v35 := []int{(^(v3)), v3, 0}
    var v36 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
