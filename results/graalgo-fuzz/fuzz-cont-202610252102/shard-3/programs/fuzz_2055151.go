package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint; F1 int16; F2 int32 }

func main() {
    var v1 uint8 = 0
    var v2 int64 = 42
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v5 float32 = 3.14159
    for fi6 := 0; fi6 < 2; fi6++ {
        v1 = (v1 & (v1 ^ v1))
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 bool = false
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 int64 = 470
    switch ("foo bar" + ("foo bar" + "")) {
    case "foo bar":
        switch 32768 {
        case 4294967295:
            if false {
                if false {
                    var v10 int16 = 514
                    _ = v10
                }
                for fi11 := 0; fi11 < 17; fi11++ {
                    var v12 int64 = (v9 / ((v9) | 1))
                    v13 := GenS1{F0: (^(v1)), F1: 423, F2: ("\n\t" + "hello")}
                    var v14 float32 = (v5 + (v5 / 0.1))
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                v15 := GenS1{F0: (v1 / ((v1) | 1)), F1: 0, F2: "a"}
                _ = v15
            } else {
                v2 = v2
                v16 := GenS1{F0: 42, F1: 2, F2: "hello"}
                v17 := GenS1{F0: v1, F1: 2, F2: "a"}
                _ = v16
                _ = v17
            }
        case 128:
            _ = len(v4)
        case 1:
            v5 += ((3.14159 - 1.5) - v5)
        }
    }
    v2 += (v2 / ((v9) | 1))
    var v18 int = int(v2)
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 float64 = 1.5
        v21 := v3[(^(fi19))]
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 int16 = -32768
    v1 += 127
    var v23 int16 = v22
    var v24 int8 = -26
    v25 := GenS2{F0: 32767, F1: 1}
    if false {
        var v26 float32 = v5
        if (!(!true)) {
            switch -1 {
            case -9223372036854775808:
                v9 += (v2 * v9)
            case -2147483648:
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 uint16 = 42
                    _ = fi27
                    _ = v28
                }
            }
            if false {
                v2 -= 890
                v23 += 0
            } else {
                v2 += (^(v2))
                fmt.Println(v1, v24, v9)
                switch "\n\t" {
                case "":
                    v29 := []int{v18, v18, (v18 | (v18 * int(v24))), (v18 << ((2) & 63))}
                    _ = v29
                case "a":
                    var v30 int8 = (v24 << ((18446744073709551615) & 63))
                    _ = v30
                case "foo bar":
                    var v31 int = v18
                    _ = v31
                default:
                    switch (v18 / ((v18) | 1)) {
                    case 4294967295:
                        v32 := GenS3{F0: (-(v23)), F1: (v24 / ((int8(v1)) | 1))}
                        _ = v32
                    case 128:
                        fmt.Println(v1)
                    case 2147483647:
                        fmt.Println(v24, v5, v24)
                    default:
                        v33 := GenS4{F0: 501, F1: v22, F2: 198}
                        _ = v33
                    }
                }
            }
            v34 := []int{}
            _ = v34
        }
        _ = v26
    } else {
        fmt.Println(v22, v5, v1)
    }
    var v35 int32 = 2147483647
    v36 := GenS3{F0: -32768, F1: v24}
    v37 := GenS1{F0: v1, F1: 371, F2: ("foo bar" + "\n\t")}
    var v38 int8 = (v24 << ((172) & 63))
    if false {
        for fi39 := 0; fi39 < 11; fi39++ {
            fmt.Println(fi39, v38, v18)
            var v40 float64 = 3.14159
            var v41 uint32 = 65535
            _ = fi39
            _ = v40
            _ = v41
        }
        var v42 int32 = v35
        _ = v42
    } else {
        v43 := GenS3{F0: v23, F1: 127}
        v44 := GenS2{F0: 2, F1: 2}
        switch "\n\t" {
        case "":
            switch (v18 | v18) {
            case -2147483648:
                var v45 int8 = -92
                _ = v45
            }
        case "hello":
            v46 := []int{v18, (v18 * (v18 ^ v18)), (v18 >> ((91) & 63))}
            _ = v46
        default:
            var v47 int8 = -12
            _ = v47
        }
        _ = v43
        _ = v44
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        v49 := v4[(fi48 & (fi48 - int(v35)))]
        if (true && (2147483647 <= v35)) {
            var v50 uint32 = 0
            _ = v50
        }
        v51 := GenS2{F0: 0, F1: 2147483648}
        _ = fi48
        _ = v49
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v9
}
