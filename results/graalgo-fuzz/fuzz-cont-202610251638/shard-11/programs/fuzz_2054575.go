package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint = 255
    var v2 float32 = -2.25
    var v3 bool = (1 <= 2)
    var v4 uint32 = 970
    var v5 float32 = ((v2 / (1e10 * v2)) * v2)
    v6 := GenS1{F0: ""}
    v1 = v1
    switch 9223372036854775807 {
    case -343:
        switch ((("a" + ("foo bar" + "")) + "hello") + "") {
        case "a":
            v7 := []int{}
            _ = v7
        case "\n\t":
            fmt.Println(v3, v4, v3)
        case "":
            v2 += 3.14159
        }
    default:
        if (-9223372036854775808 == -1) {
            switch "hello" {
            case "foo bar":
                v8 := []int{128, -9223372036854775808, 164}
                _ = v8
            case "\n\t":
                for fi9 := 0; fi9 < 10; fi9++ {
                    fmt.Println(v2, v4)
                    if true {
                        var v10 uint32 = v4
                        _ = v10
                    }
                    v3 = v3
                    _ = fi9
                }
            case "a":
                v5 -= ((v2 * v2) * (((v5 + v2) + (-2.25 * float32((1e10 * -2.25)))) - (((float32(0.1) + 0.0) + (((v2 + float32(-2.25)) / v5) / v5)) + float32((3.14159 / ((float64(v5) + -2.25) * ((-2.25 - 3.14159) / 1.0)))))))
            }
        } else {
            var v11 float32 = v5
            v11 += ((float32((-2.25 / ((1.5 + 0.0) + 1e10))) - float32(0.0)) - v5)
            _ = v11
        }
    }
    if (0 == 100) {
        switch 753 {
        case 0:
            if v3 {
                var v12 int = 42
                switch "a" {
                case "":
                    var v13 int32 = 0
                    _ = v13
                case "hello":
                    v3 = v3
                case "foo bar":
                    v3 = (!true)
                }
                v14 := GenS1{F0: "hello"}
                _ = v12
                _ = v14
            }
        case 136:
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 int32 = 0
                _ = fi15
                _ = v16
            }
        case -382:
            v17 := []int{127, 127, 128, 42}
            _ = v17
        }
        switch -1 {
        case 32767:
            if v3 {
                for fi18 := 0; fi18 < 14; fi18++ {
                    fmt.Println(fi18)
                    if v3 {
                        v19 := []int{-573, int(v4), 1, (fi18 & fi18)}
                        switch "\n\t" {
                        case "":
                            var v20 float64 = float64(v5)
                            _ = v20
                        default:
                            var v21 int8 = 0
                            _ = v21
                        }
                        _ = v19
                    }
                    if (true && v3) {
                        fmt.Println(v2, v4)
                        fmt.Println(v4, fi18, v4)
                    }
                    _ = fi18
                }
                switch ("\n\t" + "a") {
                case "a":
                    for fi22 := 0; fi22 < 5; fi22++ {
                        var v23 int8 = 42
                        switch fi22 {
                        case 128:
                            fmt.Println(v23, v1)
                        case 0:
                            v24 := GenS1{F0: "foo bar"}
                            _ = v24
                        case 32768:
                            fmt.Println(v3, fi22)
                        }
                        v25 := []int{303, -1, 32768, fi22}
                        _ = fi22
                        _ = v23
                        _ = v25
                    }
                case "foo bar":
                    v5 += v2
                case "hello":
                    v4 += (v4 + v4)
                default:
                    v26 := GenS2{F0: 4294967295}
                    _ = v26
                }
            }
        case 2147483647:
            v3 = ((v1 < v1) && v3)
        default:
            v5 -= float32(float64(((v5 - (v5 + (float32(((0.1 - 0.1) - float64(v2))) / ((v2 / float32(3.14159)) - v5)))) + (((v2 + v2) * v2) - (v5 * (v2 - (v5 - (v5 * v2))))))))
        }
    }
    v5 = (v2 * (((v5 - float32(0.0)) / (v2 * 1.5)) - v2))
    var v27 uint16 = 2
    var v28 float64 = 1.5
    var v29 uint8 = 0
    if (1 <= -2147483648) {
        var v30 int64 = 9223372036854775807
        v31 := GenS1{F0: ""}
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
