package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 string; F2 int16 }

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]bool)
    var v3 uint16 = 255
    v4 := make(map[uint32]bool)
    var v5 bool = false
    var v6 int32 = 1
    switch "hello" {
    case "foo bar":
        var v7 uint16 = v3
        _ = v7
    case "":
        var v8 int32 = v6
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 17; fi9++ {
            fmt.Println(fi9, v6)
            v5 = (444 >= (v3 << ((401) & 63)))
            fmt.Println(v1, v5, v3)
            _ = fi9
        }
    }
    v10 := v4[626]
    v1 += -2147483648
    switch (^(v1)) {
    case 9223372036854775807:
        if ((!true) || (!(!false))) {
            var v11 float64 = 0.0
            _ = v11
        } else {
            if v5 {
                for fi12 := 0; fi12 < 8; fi12++ {
                    v13 := v4[74]
                    _ = fi12
                    _ = v13
                }
                v2[65535] = v10
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 uint64 = 1
                    var v16 float32 = 0.0
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
        }
    }
    var v17 uint64 = 683
    delete(v4, 794)
    v18 := GenS1{F0: v3}
    if ((1.5 < -2.25) && (v10 && v5)) {
        if ((2 > 4294967295) && ((v3 ^ (v3 & 65535)) < (v3 - v3))) {
            var v19 bool = v5
            _ = v19
        }
        switch "foo bar" {
        case "\n\t":
            v6 = v6
        case "foo bar":
            v10 = true
        default:
            v17 -= 2
        }
    }
    switch "foo bar" {
    case "":
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := []int{fi20}
            if (1e10 < 0.0) {
                for fi22 := 0; fi22 < 6; fi22++ {
                    if v5 {
                        var v23 string = "foo bar"
                        _ = v23
                    }
                    _ = fi22
                }
                for fi24 := 0; fi24 < 3; fi24++ {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    _ = fi24
                }
                for fi25 := 0; fi25 < 5; fi25++ {
                    if true {
                        var v26 uint64 = uint64(v3)
                        v21 = append(v21, 42)
                        v2[871] = true
                        _ = v26
                    } else {
                        fmt.Println(v3)
                        v27 := GenS2{F0: 32767, F1: "a", F2: -32768}
                        fmt.Println(v5)
                        _ = v27
                    }
                    _ = fi25
                }
            }
            _ = fi20
            _ = v21
        }
    default:
        v5 = v10
    }
    if true {
        v28 := v2[2147483647]
        _ = v28
    } else {
        var v29 float32 = 1e10
        var v30 int64 = -9223372036854775808
        _ = v29
        _ = v30
    }
    if v5 {
        v31 := GenS2{F0: 42, F1: "\n\t", F2: -1}
        if (42 == (^(v6))) {
            var v32 bool = true
            fmt.Println(v10, v1, v3)
            v10 = true
            _ = v32
        } else {
            v5 = v10
            if true {
                v10 = (((3.14159 * 3.14159) + 1.5) >= float32(1e10))
                if (false || ((-131 <= -358) && v10)) {
                    if (true || (861 <= 4294967295)) {
                        v33 := []int{v1, 32767, v1, (^(v1))}
                        var v34 string = (("hello" + ("\n\t" + ("foo bar" + ""))) + "hello")
                        _ = v33
                        _ = v34
                    }
                    fmt.Println(v6, v10, v10)
                    var v35 int64 = 1
                    _ = v35
                }
            }
            var v36 float64 = float64(1.5)
            _ = v36
        }
        v17 = 473
        _ = v31
    } else {
        if (false || (!true)) {
            fmt.Println(v1)
            v3 += v3
            v37 := GenS2{F0: 32767, F1: "", F2: -1}
            _ = v37
        }
    }
    switch ("a" + "a") {
    case "hello":
        var v38 int64 = 9223372036854775807
        _ = v38
    case "a":
        v39 := GenS1{F0: v3}
        _ = v39
    default:
        v40 := GenS1{F0: 688}
        _ = v40
    }
    var v41 uint64 = (v17 / ((uint64(v3)) | 1))
    switch "\n\t" {
    case "hello":
        var v42 int64 = 80
        _ = v42
    case "foo bar":
        v17 -= v17
    }
    fmt.Println(v5, v3)
    v43 := GenS2{F0: 1000, F1: "a", F2: 514}
    if true {
        for fi44 := 0; fi44 < 10; fi44++ {
            for fi45 := 0; fi45 < 13; fi45++ {
                if v10 {
                    var v46 uint = 1
                    var v47 float32 = (-2.25 * 0.1)
                    v10 = v10
                    var v48 int16 = 206
                    var v49 uint16 = v3
                    _ = v46
                    _ = v47
                    _ = v48
                    _ = v49
                }
                _ = fi45
            }
            _ = fi44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v5
    _ = v6
}
