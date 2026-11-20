package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int64; F1 int; F2 int64 }
func genFn0() int64 {
    v7 := GenS1{F0: -1}
    var v8 float64 = (1.5 - 3.14159)
    if false {
        var v9 int64 = -9223372036854775808
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 string = "foo bar"
                fmt.Println(v9, fi10)
                v13 := GenS1{F0: (fi11 >> ((255) & 63))}
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 16; fi14++ {
                v8 = float64((float32((float64(-2.25) - v8)) - (0.1 - float32(1.5))))
                var v15 uint8 = 103
                _ = fi14
                _ = v15
            }
            _ = fi10
        }
        _ = v9
    }
    switch "foo bar" {
    case "foo bar":
        switch (("a" + "\n\t") + "hello") {
        case "hello":
            v8 = ((v8 * ((((1.5 / v8) + v8) + ((v8 - 3.14159) / (v8 + (float64(1.5) / 1.5)))) / 1.5)) * (v8 / ((0.0 / 3.14159) - ((v8 + (v8 - v8)) * ((v8 / v8) + v8)))))
        case "":
            v16 := GenS2{F0: 1, F1: 18446744073709551615, F2: 42}
            _ = v16
        case "foo bar":
            fmt.Println(v8)
        default:
            v8 = v8
        }
    case "\n\t":
        if true {
            v8 -= 1.5
            var v17 uint16 = 42
            for fi18 := 0; fi18 < 4; fi18++ {
                if ((v8 <= v8) || (true && (true && (true && false)))) {
                    fmt.Println(v17, fi18, v8)
                    v19 := GenS3{F0: 4294967295}
                    _ = v19
                }
                v20 := GenS3{F0: (fi18 ^ fi18)}
                v8 += v8
                _ = fi18
                _ = v20
            }
            _ = v17
        }
    case "a":
        v8 += float64(float32(3.14159))
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 217
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    v5 := []string{"a"}
    var v6 int16 = 32767
    var v21 uint16 = 2
    var v22 uint8 = v2
    if true {
        switch 207 {
        case 128:
            for fi23 := 0; fi23 < 8; fi23++ {
                switch "\n\t" {
                case "foo bar":
                    v2 = v22
                case "a":
                    switch "" {
                    case "":
                        var v24 float32 = 1.5
                        _ = v24
                    case "foo bar":
                        v5 = append(v5, ("\n\t" + "a"))
                    }
                default:
                    var v25 string = "hello"
                    _ = v25
                }
                _ = fi23
            }
        }
        delete(v3, 4294967295)
        if (!(3.14159 <= 1.5)) {
            switch "foo bar" {
            case "a":
                v3[127] = (false || true)
            }
        }
    } else {
        v26 := GenS4{F0: 80, F1: (v22 ^ v2)}
        var v27 int16 = v6
        var v28 int64 = -146
        _ = v26
        _ = v27
        _ = v28
    }
    if false {
        for fi29 := 0; fi29 < 12; fi29++ {
            switch ("foo bar" + "") {
            case "hello":
                var v30 string = "a"
                _ = v30
            case "foo bar":
                var v31 string = ""
                _ = v31
            case "\n\t":
                v6 = v6
            }
            _ = fi29
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            v33 := GenS5{F0: 0, F1: fi32, F2: genFn0()}
            var v34 float64 = 0.1
            _ = fi32
            _ = v33
            _ = v34
        }
        for fi35 := 0; fi35 < 7; fi35++ {
            if (52 == 218) {
                fmt.Println(v21)
                for fi36 := 0; fi36 < 2; fi36++ {
                    var v37 uint32 = 1
                    v2 = uint8(v37)
                    _ = fi36
                    _ = v37
                }
                if false {
                    var v38 bool = (369 != 0)
                    _ = v38
                }
            } else {
                var v39 uint16 = v21
                switch "\n\t" {
                case "hello":
                    delete(v1, 2)
                default:
                    if (v39 < (v21 - (v39 << ((0) & 63)))) {
                        v40 := GenS3{F0: fi35}
                        _ = v40
                    } else {
                        var v41 bool = false
                        var v42 uint64 = 454
                        v2 += v2
                        _ = v41
                        _ = v42
                    }
                }
                _ = v39
            }
            var v43 uint = 18446744073709551615
            _ = fi35
            _ = v43
        }
    }
    fmt.Println(v6)
    var v44 string = "foo bar"
    fmt.Println(v21, v44)
    for fi45 := 0; fi45 < 5; fi45++ {
        var v46 int64 = -783
        _ = fi45
        _ = v46
    }
    v3[128] = (false || true)
    var v47 bool = (!false)
    var v48 bool = v47
    switch 861 {
    case 4294967295:
        var v49 float64 = 3.14159
        _ = v49
    default:
        _ = len(v5)
    }
    v50 := GenS1{F0: int(v22)}
    var v51 bool = false
    v5 = append(v5, v44)
    var v52 float32 = 3.14159
    v53 := GenS1{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
