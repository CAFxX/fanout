package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int32 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint32; F2 int32 }
func genFn0(p0 uint16) string {
    p0 -= 2
    for fi5 := 0; fi5 < 11; fi5++ {
        p0 = (p0 - p0)
        if (!true) {
            v6 := GenS1{F0: 2, F1: 682, F2: -2147483648}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 int8 = -128
                v9 := GenS1{F0: 42, F1: 22, F2: -1}
                _ = fi7
                _ = v8
                _ = v9
            }
            v10 := []int{754, (fi5 ^ (fi5 - (fi5 + (-(fi5))))), fi5}
            var v11 uint8 = 0
            _ = v10
            _ = v11
        }
        var v12 uint8 = 99
        _ = fi5
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        if false {
            var v14 float64 = (1e10 - (((((((0.1 * (1.5 - 0.1)) / 1.0) / 1.5) + float64(float32(1e10))) + (0.0 * (float64(0.1) * float64(float32(3.14159))))) / (3.14159 / (1e10 * 1.5))) / 1e10))
            v15 := []int{}
            var v16 uint16 = (p0 ^ (^(p0)))
            _ = v14
            _ = v15
            _ = v16
        } else {
            if true {
                fmt.Println(p0, p0)
                var v17 int16 = 0
                _ = v17
            } else {
                p0 = p0
                v18 := []int{fi13}
                _ = v18
            }
            var v19 uint = 42
            var v20 float32 = ((1.5 - (0.0 + 1.5)) * ((((3.14159 + ((float32(0.0) + ((float32(1e10) * float32((-2.25 * 1.5))) - 1.5)) * -2.25)) / 0.1) - 1.5) + 0.0))
            _ = v19
            _ = v20
        }
        _ = fi13
    }
    var v21 bool = (-1 == 0)
    switch ("foo bar" + "") {
    case "":
        for fi22 := 0; fi22 < 20; fi22++ {
            switch fi22 {
            case 0:
                for fi23 := 0; fi23 < 4; fi23++ {
                    if v21 {
                        v24 := []int{}
                        fmt.Println(fi22, v21, fi23)
                        v25 := GenS2{F0: "foo bar", F1: ("a" + "")}
                        _ = v24
                        _ = v25
                    } else {
                        v26 := GenS3{F0: v21, F1: 255}
                        var v27 bool = v21
                        var v28 string = ""
                        var v29 uint8 = 42
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    var v30 uint16 = p0
                    var v31 uint64 = 18446744073709551615
                    _ = fi23
                    _ = v30
                    _ = v31
                }
            }
            _ = fi22
        }
    case "foo bar":
        v32 := GenS4{F0: 65535}
        _ = v32
    case "hello":
        p0 = (p0 + 0)
    default:
        if v21 {
            fmt.Println(p0)
            v21 = (2 >= 42)
        }
    }
    _ = p0
    _ = v21
    return "\n\t"
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    var v4 uint8 = 255
    switch "" {
    case "":
        v33 := GenS5{F0: 2147483647, F1: 2147483647, F2: 0}
        _ = v33
    default:
        v34 := []int{1}
        _ = v34
    }
    v2[65535] = 0
    v35 := GenS4{F0: 2}
    var v36 string = ((("\n\t" + "hello") + "a") + "")
    for fi37 := 0; fi37 < 13; fi37++ {
        v2[622] = 856
        switch ("\n\t" + "hello") {
        case "\n\t":
            v2[737] = 2
        case "foo bar":
            v3 -= v3
        default:
            for fi38 := 0; fi38 < 2; fi38++ {
                var v39 string = v36
                var v40 float64 = 0.1
                _ = fi38
                _ = v39
                _ = v40
            }
        }
        _ = fi37
    }
    v3 += 1
    _ = len(v2)
    for fi41 := 0; fi41 < 11; fi41++ {
        v36 = "foo bar"
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
