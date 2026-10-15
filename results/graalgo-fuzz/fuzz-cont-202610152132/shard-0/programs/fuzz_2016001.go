package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int32 }

func main() {
    var v1 float32 = (float32(0.1) - ((float32(1.5) * -2.25) + -2.25))
    var v2 float32 = (v1 + v1)
    v3 := []string{"\n\t", "\n\t", ""}
    v2 += -2.25
    v2 = (v2 * 1e10)
    switch -663 {
    case 0:
        fmt.Println(v2, v2)
    case 487:
        for fi4 := 0; fi4 < 13; fi4++ {
            v2 = v2
            _ = fi4
        }
    default:
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 uint = 18446744073709551615
            if (1.5 >= 1e10) {
                var v7 float64 = 1.5
                _ = v7
            } else {
                v8 := GenS1{F0: 0, F1: fi5}
                _ = v8
            }
            var v9 uint16 = 42
            _ = fi5
            _ = v6
            _ = v9
        }
    }
    v1 = 1e10
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int = fi10
            fmt.Println(v2, fi10, fi10)
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 float32 = (v2 * (v2 + -2.25))
        _ = fi13
        _ = v14
    }
    var v15 int8 = -1
    switch "\n\t" {
    case "":
        v1 -= (-2.25 / v1)
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = "a"
        }
    case "\n\t":
        fmt.Println(v15)
    }
    v15 = (v15 - v15)
    v3 = append(v3, ("a" + ""))
    for fi16 := 0; fi16 < 4; fi16++ {
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v17 := GenS1{F0: 558, F1: 127}
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS2{F0: fi18}
            var v20 int8 = int8(fi18)
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: 0, F1: fi16}
        _ = fi16
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint8 = 2
        for fi24 := 0; fi24 < 18; fi24++ {
            v15 = v15
            v15 = v15
            if true {
                if ((true || ("a" == ("hello" + ("foo bar" + "foo bar")))) && (42 <= 1)) {
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                    _ = len(v3)
                }
                var v25 bool = (!true)
                v26 := GenS2{F0: (fi24 / (((fi24 * (fi24 + fi22))) | 1))}
                _ = v25
                _ = v26
            }
            _ = fi24
        }
        switch -2147483648 {
        case 127:
            if len(v3) > 0 {
                v3[0] = (("\n\t" + "foo bar") + "foo bar")
            }
        case -555:
            if false {
                v15 += (v15 | (v15 | (^(v15))))
                v15 += (v15 | (v15 >> ((798) & 63)))
            } else {
                fmt.Println(v15)
            }
        }
        _ = fi22
        _ = v23
    }
    var v27 string = "a"
    v28 := GenS3{F0: 0}
    switch 42 {
    case -1:
        v29 := GenS4{F0: 65535, F1: 0, F2: -2147483648}
        _ = v29
    case 4294967295:
        switch -9223372036854775808 {
        case -2147483648:
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 uint16 = 237
                if false {
                    var v32 int = (fi30 + fi30)
                    switch v27 {
                    case "\n\t":
                        v1 -= ((0.1 - float32(0.0)) / v2)
                    case "hello":
                        var v33 string = ("foo bar" + (v27 + v27))
                        _ = v33
                    case "foo bar":
                        var v34 string = v27
                        _ = v34
                    }
                    var v35 int16 = 32767
                    _ = v32
                    _ = v35
                }
                _ = fi30
                _ = v31
            }
        case -1:
            switch (v27 + ((("" + "hello") + "a") + (v27 + (((v27 + (v27 + ((("foo bar" + (v27 + ((((((v27 + "foo bar") + (v27 + "a")) + v27) + "a") + (v27 + (v27 + "a"))) + "\n\t"))) + v27) + ""))) + ("a" + v27)) + ((v27 + "a") + v27))))) {
            case "hello":
                var v36 float64 = float64(1.5)
                _ = v36
            case "":
                fmt.Println(v15, v15)
            }
        case 128:
            if true {
                var v37 int64 = 1
                _ = v37
            } else {
                v3 = append(v3, v27)
                fmt.Println(v2)
                v15 += (v15 ^ (v15 & -128))
            }
        default:
            _ = len(v3)
        }
    case -268:
        v38 := GenS1{F0: 65535, F1: 377}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
