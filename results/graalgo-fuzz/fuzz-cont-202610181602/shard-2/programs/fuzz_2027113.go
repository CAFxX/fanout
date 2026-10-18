package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int64; F2 uint8 }
func genFn0(p0 string, p1 uint) uint16 {
    if (18446744073709551615 < 42) {
        for fi5 := 0; fi5 < 1; fi5++ {
            var v6 uint8 = 63
            v7 := GenS1{F0: (p1 ^ (p1 + 612))}
            var v8 uint64 = 0
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        v9 := []int{}
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int = (^(fi10))
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint = p1
            switch v11 {
            case 1:
                var v14 uint64 = 412
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        switch (v11 >> ((1) & 63)) {
        case -870:
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 float64 = 1.5
                if (-831 < 42) {
                    var v17 int64 = -150
                    _ = v17
                } else {
                    v18 := GenS2{F0: -640}
                    var v19 float32 = 1.5
                    p0 = p0
                    if (true && false) {
                        v20 := []int{(v11 / ((fi10) | 1)), fi10, v11}
                        v21 := GenS2{F0: 32767}
                        _ = v20
                        _ = v21
                    } else {
                        if (((!(((1 < 0) || (140 == 32767)) && (-438 < 1))) && (((fi10 + 9223372036854775807) < (v11 - (-(fi10)))) || false)) || (p1 != p1)) {
                            var v22 uint16 = 827
                            _ = v22
                        }
                        var v23 int = v11
                        _ = v23
                    }
                    v24 := []int{v11}
                    _ = v18
                    _ = v19
                    _ = v24
                }
                fmt.Println(v16, fi15, v16)
                var v25 uint8 = 127
                _ = fi15
                _ = v16
                _ = v25
            }
        default:
            v26 := []int{fi10, 127, v11}
            _ = v26
        }
        _ = fi10
        _ = v11
    }
    if false {
        var v27 uint16 = 42
        v28 := GenS1{F0: (p1 << ((p1) & 63))}
        _ = v27
        _ = v28
    }
    switch "a" {
    case "a":
        var v29 uint16 = 32767
        _ = v29
    case "hello":
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 uint32 = 65535
            v32 := GenS3{F0: 18446744073709551615}
            var v33 int32 = 0
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
    case "foo bar":
        var v34 uint16 = 2
        _ = v34
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int16 = 1
    var v2 int32 = 2147483647
    var v3 int64 = 42
    var v4 float32 = 1.5
    v35 := GenS2{F0: int16(v2)}
    switch 997 {
    case -2147483648:
        var v36 uint16 = 1
        _ = v36
    case 42:
        v37 := GenS2{F0: (v1 >> ((989) & 63))}
        _ = v37
    case 919:
        v38 := GenS3{F0: 42}
        _ = v38
    }
    v2 -= v2
    var v39 int16 = (v1 ^ 394)
    var v40 float64 = 3.14159
    v3 += v3
    for fi41 := 0; fi41 < 7; fi41++ {
        v42 := GenS4{F0: "foo bar"}
        _ = fi41
        _ = v42
    }
    v43 := []int{}
    var v44 float64 = v40
    if (("\n\t" + "a") > "hello") {
        v43 = append(v43, 2147483647)
        var v45 int32 = 4
        v3 -= (v3 | (v3 / ((v3) | 1)))
        _ = v45
    } else {
        v46 := GenS4{F0: "\n\t"}
        var v47 int8 = 31
        if false {
            v44 = (float64(v4) * float64((1e10 / -2.25)))
        }
        _ = v46
        _ = v47
    }
    v1 = 32767
    for fi48 := 0; fi48 < 5; fi48++ {
        if len(v43) > 0 {
            v43[0] = (fi48 / (((fi48 << ((255) & 63))) | 1))
        }
        for fi49 := 0; fi49 < 1; fi49++ {
            fmt.Println(v4, v3, v4)
            var v50 string = "a"
            _ = fi49
            _ = v50
        }
        switch "" {
        case "":
            v51 := GenS4{F0: "\n\t"}
            _ = v51
        case "foo bar":
            v52 := GenS5{F0: "", F1: 1, F2: 1}
            _ = v52
        }
        _ = fi48
    }
    v53 := GenS2{F0: -298}
    var v54 int64 = 0
    v44 = v40
    if len(v43) > 0 {
        v43[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v53
    _ = v54
}
