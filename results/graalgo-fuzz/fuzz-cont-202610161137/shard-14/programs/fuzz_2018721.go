package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint32) int32 {
    switch 128 {
    case -2147483648:
        p0 += p0
    case 42:
        for fi4 := 0; fi4 < 18; fi4++ {
            var v5 uint8 = 156
            if false {
                for fi6 := 0; fi6 < 10; fi6++ {
                    switch "a" {
                    case "":
                        v7 := []int{fi4}
                        _ = v7
                    }
                    v8 := []int{0, (fi4 + fi4)}
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi6
                    _ = v8
                }
                v9 := GenS1{F0: p0, F1: 42, F2: v5}
                _ = v9
            }
            _ = fi4
            _ = v5
        }
    case -9223372036854775808:
        v10 := []int{-703}
        _ = v10
    }
    switch ((("foo bar" + "hello") + "hello") + "\n\t") {
    case "\n\t":
        v11 := GenS1{F0: p0, F1: 22, F2: 41}
        _ = v11
    }
    v12 := []int{-2147483648, 1, 127}
    var v13 uint8 = 127
    _ = p0
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 int = -2147483648
    var v2 int32 = 56
    v3 := []string{"a"}
    v14 := GenS1{F0: 468, F1: 310, F2: 0}
    _ = len(v3)
    switch v1 {
    case -6:
        switch 42 {
        case 127:
            fmt.Println(v1, v1)
        case 4294967295:
            v3 = append(v3, "a")
        case 9223372036854775807:
            var v15 uint8 = 2
            _ = v15
        default:
            v1 = 4294967295
        }
    }
    v2 -= (v2 * (v2 / (((-(v2))) | 1)))
    var v16 int16 = -1
    if (!(444 == 65535)) {
        fmt.Println(v2)
        switch ("" + "a") {
        case "a":
            fmt.Println(v16, v16, v16)
        case "hello":
            v1 -= -828
        }
        switch (v1 >> ((42) & 63)) {
        case -317:
            _ = len(v3)
        }
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS1{F0: 2147483647, F1: 152, F2: 167}
        fmt.Println(v2, v1, fi17)
        v16 = (v16 % ((v16) | 1))
        _ = fi17
        _ = v18
    }
    switch ("\n\t" + ("foo bar" + "")) {
    case "foo bar":
        fmt.Println(v1, v16, v16)
    case "\n\t":
        for fi19 := 0; fi19 < 14; fi19++ {
            switch fi19 {
            case 630:
                fmt.Println(v2, v2)
            case 731:
                v1 = (-(fi19))
            case 35:
                v1 = v1
            }
            _ = fi19
        }
    }
    if (!(!false)) {
        v16 -= -301
        var v20 float32 = (((1e10 * 0.1) / 1e10) / 3.14159)
        _ = v20
    } else {
        if false {
            v1 = -2147483648
            var v21 uint32 = 2147483647
            if (118 == 90) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v22 := GenS2{F0: 2, F1: -32768}
                _ = v22
            } else {
                var v23 int16 = v16
                v16 = v23
                fmt.Println(v1, v1, v2)
                _ = v23
            }
            _ = v21
        }
    }
    var v24 float64 = 0.1
    fmt.Println(v2, v1, v1)
    for fi25 := 0; fi25 < 15; fi25++ {
        _ = len(v3)
        _ = fi25
    }
    v16 -= 32767
    v16 += 0
    fmt.Println(v1)
    var v26 float64 = (v24 / v24)
    v27 := GenS1{F0: 1, F1: 42, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
}
