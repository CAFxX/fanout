package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint64; F1 int32 }
func genFn0(p0 bool) int32 {
    p0 = (203 < 0)
    switch ("\n\t" + "a") {
    case "\n\t":
        v6 := []int{-9223372036854775808, -2147483648, -758, 42}
        _ = v6
    case "hello":
        v7 := GenS1{F0: ("foo bar" + ""), F1: 128, F2: 0}
        _ = v7
    }
    _ = p0
    return 633
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1e10
    v3 := make(map[int]int)
    var v4 int8 = 52
    var v5 uint64 = 0
    v8 := GenS2{F0: true}
    fmt.Println(v2, v5)
    var v9 uint64 = v5
    v10 := GenS2{F0: true}
    v5 = v9
    v11 := []int{128, -954, 954, -171}
    fmt.Println(v9, v2)
    var v12 int16 = 1000
    switch 426 {
    case -667:
        v5 = (v9 & 4294967295)
    case 154:
        var v13 int16 = 42
        _ = v13
    case -2147483648:
        fmt.Println(v5, v2, v4)
    default:
        fmt.Println(v2)
    }
    v14 := GenS3{F0: 4294967295, F1: 0, F2: (v12 / (((^(v12))) | 1))}
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            switch 0 {
            case 2147483647:
                v17 := GenS3{F0: (v9 | (v5 + 2)), F1: 170, F2: (v12 * -32768)}
                _ = v17
            case 4294967295:
                var v18 float32 = v2
                _ = v18
            }
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            switch (fi15 >> ((4294967295) & 63)) {
            case 291:
                var v19 string = (("" + "hello") + "foo bar")
                _ = v19
            }
            _ = fi16
        }
        for fi20 := 0; fi20 < 14; fi20++ {
            switch "a" {
            case "hello":
                if (42 == int32(v9)) {
                    v9 = 950
                    var v21 uint16 = 899
                    v22 := GenS2{F0: (!("" <= ""))}
                    if (true || (4294967295 >= 2147483647)) {
                        fmt.Println(v5, v12)
                    } else {
                        var v23 uint = 4294967295
                        v24 := GenS1{F0: "a", F1: 255, F2: int8(v21)}
                        _ = v23
                        _ = v24
                    }
                    v25 := v1[fi20]
                    _ = v21
                    _ = v22
                    _ = v25
                }
            }
            if (float32((float64(float32((0.1 * 0.1))) * 0.1)) >= 1e10) {
                v26 := GenS3{F0: 153, F1: 42, F2: (^(v12))}
                var v27 uint32 = 177
                _ = v26
                _ = v27
            } else {
                var v28 uint64 = 0
                _ = v28
            }
            _ = fi20
        }
        _ = fi15
    }
    var v29 uint32 = 191
    var v30 int32 = -915
    for fi31 := 0; fi31 < 17; fi31++ {
        if (!true) {
            if false {
                switch (-(fi31)) {
                case 4294967295:
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 uint32 = 4294967295
                        var v34 bool = true
                        v35 := GenS3{F0: (v9 & v9), F1: 175, F2: v12}
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
                v36 := GenS3{F0: v5, F1: 0, F2: v12}
                _ = v36
            }
            v37 := GenS4{F0: (v9 % (((v9 + 2)) | 1)), F1: genFn0(false)}
            _ = v37
        }
        _ = fi31
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 uint8 = 2
        fmt.Println(v30, v12, v5)
        _ = fi38
        _ = v39
    }
    v2 += (v2 / 1e10)
    for fi40 := 0; fi40 < 4; fi40++ {
        if false {
            v41 := GenS1{F0: "", F1: 127, F2: (v4 | (v4 % ((-117) | 1)))}
            v29 -= (v29 ^ v29)
            _ = v41
        } else {
            var v42 int16 = v12
            v43 := v1[fi40]
            _ = v42
            _ = v43
        }
        fmt.Println(v30)
        v44 := GenS2{F0: (0 < 2)}
        _ = fi40
        _ = v44
    }
    if false {
        v3[427] = 549
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
