package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int32; F1 bool; F2 int8 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (-2.25 - -2.25)
    var v3 bool = (true || (42 <= -1))
    v4 := make(map[int]bool)
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: ("" + ("foo bar" + ""))}
            fmt.Println(fi6)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    fmt.Println(v3)
    v3 = (("a" + ("\n\t" + "\n\t")) >= "hello")
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int = -2147483648
            _ = fi9
            _ = v10
        }
        v3 = false
        var v11 int = fi8
        _ = fi8
        _ = v11
    }
    v3 = v3
    fmt.Println(v3, v2)
    for fi12 := 0; fi12 < 17; fi12++ {
        v2 = 0.0
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 string = "\n\t"
            var v15 int64 = -9223372036854775808
            v16 := GenS1{F0: (v14 + v14)}
            v14 = v14
            v17 := GenS1{F0: v14}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
            _ = v17
        }
        var v18 int64 = -9223372036854775808
        _ = fi12
        _ = v18
    }
    v19 := []int{}
    if ((42 != 254) && v3) {
        if v3 {
            v3 = (!v3)
        } else {
            fmt.Println(v2, v3)
            fmt.Println(v2, v2)
            v19 = append(v19, -1)
        }
        fmt.Println(v2, v2)
    } else {
        v20 := GenS1{F0: "hello"}
        v21 := GenS2{F0: "foo bar", F1: 65535}
        v3 = (v3 || ("hello" >= "hello"))
        _ = v20
        _ = v21
    }
    v2 = 1.5
    v3 = v3
    for fi22 := 0; fi22 < 18; fi22++ {
        v2 = v2
        _ = fi22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := GenS1{F0: "a"}
        if (!v3) {
            switch ("foo bar" + "") {
            case "hello":
                var v25 int = (fi23 / ((-867) | 1))
                _ = v25
            default:
                if (!false) {
                    v26 := GenS3{F0: -1, F1: v3, F2: int8(fi23)}
                    _ = v26
                } else {
                    _ = len(v19)
                }
            }
            if (0 == 4294967295) {
                fmt.Println(v3, v2, v3)
                switch 128 {
                case -625:
                    fmt.Println(fi23, v2, v3)
                case 4294967295:
                    var v27 int64 = 42
                    _ = v27
                }
            }
            v28 := GenS3{F0: 1, F1: (!(1000 != -91)), F2: 1}
            _ = v28
        }
        var v29 string = (("a" + (("" + "") + "foo bar")) + "")
        _ = fi23
        _ = v24
        _ = v29
    }
    var v30 uint16 = 2
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS4{F0: 980}
        switch "\n\t" {
        case "\n\t":
            for fi33 := 0; fi33 < 9; fi33++ {
                var v34 int8 = 1
                _ = fi33
                _ = v34
            }
        case "":
            fmt.Println(v3)
        default:
            if len(v19) > 0 {
                v19[0] = fi31
            }
        }
        _ = fi31
        _ = v32
    }
    var v35 uint32 = 453
    if len(v19) > 0 {
        v19[0] = -123
    }
    v35 = (v35 / ((v35) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v4
}
