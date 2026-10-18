package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int64; F2 bool }
func genFn0(p0 int32, p1 int32, p2 uint) int32 {
    if false {
        p0 = 658
    }
    if false {
        if (!(-128 == 100)) {
            for fi4 := 0; fi4 < 20; fi4++ {
                fmt.Println(p0)
                _ = fi4
            }
        }
    }
    fmt.Println(p2, p1, p0)
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := []int{fi5, -257, fi5}
        var v7 int32 = p1
        _ = fi5
        _ = v6
        _ = v7
    }
    if (2 > 992) {
        v8 := GenS1{F0: "\n\t"}
        for fi9 := 0; fi9 < 5; fi9++ {
            p2 -= 255
            _ = fi9
        }
        var v10 int32 = (p0 + (p1 >> (((p2 + p2)) & 63)))
        _ = v10
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return (^(p1))
}

func main() {
    var v1 int64 = -1
    v2 := []string{}
    v3 := make(map[string]int)
    if (0 < 42) {
        switch 42 {
        case 128:
            var v11 float32 = 1.5
            _ = v11
        case -942:
            var v12 uint16 = 546
            _ = v12
        }
    }
    v1 = 956
    var v13 float32 = 1e10
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    switch "" {
    case "a":
        _ = len(v3)
    }
    if ("\n\t" == "a") {
        _ = len(v2)
        v14 := GenS2{F0: 488}
        for fi15 := 0; fi15 < 3; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                _ = len(v3)
                _ = fi16
            }
            var v17 int = -2147483648
            fmt.Println(v1, v1, v1)
            _ = fi15
            _ = v17
        }
        _ = v14
    } else {
        var v18 uint = 884
        _ = v18
    }
    if false {
        var v19 int16 = 175
        var v20 float32 = float32(float64(-2.25))
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            delete(v3, "")
            var v22 float64 = (1.5 - 0.0)
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 float32 = v13
            fmt.Println(v1, v24)
            fmt.Println(fi23)
            _ = fi23
            _ = v24
        }
    }
    v3[("\n\t" + "a")] = -739
    for fi25 := 0; fi25 < 5; fi25++ {
        fmt.Println(fi25, v1)
        var v26 int32 = -180
        var v27 float64 = 3.14159
        _ = fi25
        _ = v26
        _ = v27
    }
    switch -813 {
    case 32767:
        for fi28 := 0; fi28 < 5; fi28++ {
            delete(v3, ("a" + "a"))
            fmt.Println(v1, v1, fi28)
            v29 := GenS2{F0: 65535}
            _ = fi28
            _ = v29
        }
    }
    var v30 uint = 391
    switch -1 {
    case 32768:
        _ = len(v3)
    default:
        v13 -= (v13 / (((v13 * v13) + v13) - float32(-2.25)))
    }
    if false {
        switch 42 {
        case 944:
            for fi31 := 0; fi31 < 1; fi31++ {
                for fi32 := 0; fi32 < 13; fi32++ {
                    switch fi31 {
                    case -279:
                        v30 = v30
                    case -523:
                        v1 = v1
                        v33 := GenS3{F0: 42, F1: v1, F2: true}
                        var v34 uint32 = 4294967295
                        _ = v33
                        _ = v34
                    }
                    _ = fi32
                }
                var v35 string = "\n\t"
                _ = fi31
                _ = v35
            }
        case 127:
            var v36 uint = v30
            _ = v36
        case -2147483648:
            if true {
                _ = len(v3)
            }
        default:
            v30 += v30
        }
        var v37 float64 = 0.1
        _ = v37
    }
    if true {
        if (false || true) {
            v30 = 2
            v38 := GenS1{F0: ((("\n\t" + "") + ("foo bar" + ("hello" + "foo bar"))) + "foo bar")}
            _ = v38
        } else {
            var v39 uint16 = 32767
            _ = v39
        }
        fmt.Println(v13, v1, v13)
    } else {
        v1 = v1
    }
    fmt.Println(v30, v30)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
}
