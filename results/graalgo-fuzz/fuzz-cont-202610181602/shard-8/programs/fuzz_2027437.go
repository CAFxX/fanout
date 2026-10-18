package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 uint16; F1 uint16; F2 bool }
type GenS4 struct { F0 int32 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 1; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            if false {
                var v6 uint8 = 30
                var v7 float32 = float32(1e10)
                var v8 int8 = -1
                _ = v6
                _ = v7
                _ = v8
            }
            var v9 string = (("a" + "hello") + (("\n\t" + ("\n\t" + ("" + ("\n\t" + "foo bar")))) + ("\n\t" + "hello")))
            v10 := GenS1{F0: fi5}
            _ = fi5
            _ = v10
            _ = v9
        }
        v11 := []int{fi4}
        v11 = append(v11, 1)
        _ = fi4
        _ = v11
    }
    switch "foo bar" {
    case "":
        v12 := []int{119, 9223372036854775807}
        _ = v12
    default:
        switch "a" {
        case "foo bar":
            switch 42 {
            case -9223372036854775808:
                if (!false) {
                    var v13 uint32 = 2147483647
                    var v14 uint32 = 1
                    var v15 int16 = 32767
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    if (255 < 4294967295) {
                        v16 := []int{127}
                        v16 = append(v16, -2147483648)
                        _ = v16
                    } else {
                        v17 := GenS1{F0: -1}
                        v18 := []int{-1, 32767, 2147483647, 4294967295}
                        var v19 int32 = -2147483648
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
            case -1:
                if (!(0 != 78)) {
                    v20 := GenS1{F0: 528}
                    _ = v20
                } else {
                    var v21 uint32 = 2
                    _ = v21
                }
            case -2147483648:
                var v22 int32 = 1
                _ = v22
            }
        }
    }
    var v23 int16 = 552
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS2{F0: false, F1: 65535, F2: "a"}
        fmt.Println(fi24, v23)
        _ = fi24
        _ = v25
    }
    fmt.Println(v23)
    _ = v23
    return 0
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]int)
    var v3 int16 = 228
    v26 := GenS3{F0: 0, F1: 65535, F2: (v3 != -1)}
    for fi27 := 0; fi27 < 15; fi27++ {
        v1 += 2147483648
        switch (fi27 * fi27) {
        case -2147483648:
            v3 += v3
        case 128:
            fmt.Println(v3)
        case 126:
            fmt.Println(v3, fi27, v1)
        }
        v1 = v1
        _ = fi27
    }
    var v28 string = "\n\t"
    fmt.Println(v28)
    for fi29 := 0; fi29 < 12; fi29++ {
        if true {
            if (!false) {
                for fi30 := 0; fi30 < 15; fi30++ {
                    v1 = 2147483647
                    var v31 float32 = -2.25
                    _ = len(v2)
                    _ = fi30
                    _ = v31
                }
            } else {
                fmt.Println(v1, fi29, fi29)
                fmt.Println(fi29)
            }
            var v32 int8 = 127
            var v33 int = fi29
            _ = v32
            _ = v33
        } else {
            if (!false) {
                v3 += (v3 % (((v3 * v3)) | 1))
                v34 := []int{}
                _ = v34
            } else {
                if false {
                    v3 = int16(fi29)
                    fmt.Println(v28, v28)
                } else {
                    var v35 int32 = -1
                    _ = v35
                }
                var v36 uint = 0
                v37 := GenS2{F0: false, F1: 559, F2: v28}
                _ = v36
                _ = v37
            }
            switch fi29 {
            case 120:
                if (255 >= 2) {
                    if (!(float64(1.5) > ((-2.25 + (((-2.25 / (0.0 + (3.14159 * 1.5))) * 1e10) + 0.1)) - 0.0))) {
                        var v38 int = (fi29 / ((fi29) | 1))
                        _ = v38
                    }
                    v28 = v28
                }
            default:
                var v39 int16 = 1000
                _ = v39
            }
        }
        v40 := GenS2{F0: true, F1: 42, F2: ""}
        _ = fi29
        _ = v40
    }
    v28 = (v28 + "")
    if true {
        if true {
            v28 = ("foo bar" + v28)
            var v41 float32 = 1.5
            var v42 int = 127
            _ = v41
            _ = v42
        }
        for fi43 := 0; fi43 < 9; fi43++ {
            if true {
                var v44 int32 = -35
                fmt.Println(v28, v28)
                _ = v44
            }
            switch "hello" {
            case "":
                var v45 int = 128
                _ = v45
            case "foo bar":
                var v46 int = fi43
                _ = v46
            case "\n\t":
                var v47 int16 = v3
                _ = v47
            }
            var v48 int64 = 302
            _ = fi43
            _ = v48
        }
        if false {
            v1 -= 2147483648
            var v49 bool = false
            _ = v49
        }
    }
    for fi50 := 0; fi50 < 2; fi50++ {
        delete(v2, v28)
        switch "hello" {
        case "foo bar":
            v51 := GenS1{F0: (fi50 % ((fi50) | 1))}
            _ = v51
        case "":
            v52 := GenS1{F0: -728}
            _ = v52
        }
        _ = fi50
    }
    for fi53 := 0; fi53 < 11; fi53++ {
        var v54 int16 = (v3 & v3)
        _ = fi53
        _ = v54
    }
    if (2147483647 < 42) {
        v55 := GenS4{F0: 0}
        _ = v55
    } else {
        for fi56 := 0; fi56 < 2; fi56++ {
            v1 = v1
            fmt.Println(v28, v28)
            v1 = (v1 + v1)
            _ = fi56
        }
        v1 = v1
    }
    switch v28 {
    case "a":
        var v57 int16 = v3
        _ = v57
    case "":
        var v58 int16 = (v3 ^ v3)
        _ = v58
    default:
        var v59 uint = 4294967295
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
}
