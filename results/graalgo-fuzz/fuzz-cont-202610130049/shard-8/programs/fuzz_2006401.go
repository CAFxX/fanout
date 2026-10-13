package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
func genFn0(p0 uint32, p1 int8, p2 string) uint {
    for fi5 := 0; fi5 < 13; fi5++ {
        p2 = p2
        var v6 uint32 = 323
        var v7 string = p2
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 698
}
func genFn1(p0 int8) uint32 {
    var v9 int = 2147483647
    fmt.Println(v9, p0)
    _ = p0
    _ = v9
    return 665
}

func main() {
    var v1 int32 = -2147483648
    var v2 int32 = (v1 >> ((65535) & 63))
    var v3 uint8 = 1
    var v4 int32 = 0
    if (((3.14159 == ((float32((0.1 * (-2.25 - 0.0))) / (1.5 - (float32(((((-2.25 - float64(0.0)) + 1e10) / (((float64(-2.25) * (float64(((-2.25 * 3.14159) * -2.25)) - -2.25)) + 1e10) - 3.14159)) + 3.14159)) + ((1.5 * float32(0.1)) / 1.0)))) - float32(((3.14159 - 1.5) + 1e10)))) || false) || false) {
        if (false && (true || (!false))) {
            switch "foo bar" {
            case "a":
                fmt.Println(v1)
            case "":
                for fi10 := 0; fi10 < 2; fi10++ {
                    fmt.Println(v4, v2, v3)
                    for fi11 := 0; fi11 < 3; fi11++ {
                        fmt.Println(fi11, v1, v4)
                        var v12 int64 = -9223372036854775808
                        if (-870 > -32768) {
                            v13 := []int{(fi10 * 1), (fi10 + 2147483647), (fi10 >> ((1) & 63))}
                            _ = v13
                        }
                        var v14 float32 = 0.1
                        if true {
                            v15 := GenS1{F0: 211, F1: (^(v12))}
                            v16 := GenS1{F0: 1000, F1: (v12 ^ v12)}
                            v17 := GenS1{F0: 1000, F1: v12}
                            _ = v15
                            _ = v16
                            _ = v17
                        } else {
                            v18 := []int{(fi10 | fi11), 2147483647, (fi10 * fi10)}
                            var v19 int = fi11
                            _ = v18
                            _ = v19
                        }
                        v20 := []int{fi10, (fi11 % ((int(v3)) | 1))}
                        _ = fi11
                        _ = v12
                        _ = v14
                        _ = v20
                    }
                    _ = fi10
                }
            case "hello":
                v21 := []int{-657, -465, 64}
                _ = v21
            }
        }
        fmt.Println(v3, v2)
        fmt.Println(v1, v3, v4)
    }
    var v22 uint64 = 169
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS2{F0: (fi23 / (((fi23 << ((4294967295) & 63))) | 1))}
        switch fi23 {
        case 657:
            fmt.Println(v4)
        case -9223372036854775808:
            v4 = v4
        case 17:
            var v25 string = "\n\t"
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v4, v3)
    v26 := GenS1{F0: 0, F1: 1}
    switch -9223372036854775808 {
    case 42:
        for fi27 := 0; fi27 < 11; fi27++ {
            if true {
                switch ("foo bar" + ("hello" + "a")) {
                case "hello":
                    if false {
                        v28 := []int{fi27, 127, 4294967295, 9223372036854775807}
                        for fi29 := 0; fi29 < 3; fi29++ {
                            v30 := GenS2{F0: (fi27 | fi29)}
                            switch "" {
                            case "a":
                                var v31 bool = (("foo bar" + (("a" + "\n\t") + "foo bar")) >= "\n\t")
                                _ = v31
                            case "hello":
                                v4 = v4
                            }
                            switch "\n\t" {
                            case "hello":
                                if true {
                                    var v32 int16 = 42
                                    _ = v32
                                }
                            }
                            _ = fi29
                            _ = v30
                        }
                        v4 += v2
                        _ = v28
                    }
                }
                fmt.Println(v3)
            } else {
                var v33 int16 = 42
                var v34 uint = 4294967295
                _ = v33
                _ = v34
            }
            for fi35 := 0; fi35 < 3; fi35++ {
                v36 := []int{(fi27 << ((2) & 63)), fi27, fi27, fi35}
                _ = fi35
                _ = v36
            }
            var v37 uint = 512
            _ = fi27
            _ = v37
        }
    case 511:
        switch "\n\t" {
        case "\n\t":
            v1 = v1
        case "a":
            v1 = (v1 ^ (v2 - v2))
        }
    default:
        for fi38 := 0; fi38 < 15; fi38++ {
            v39 := GenS3{F0: v3, F1: -121, F2: v3}
            v40 := GenS2{F0: fi38}
            var v41 int16 = -960
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
    }
    v4 += 0
    if (((float32(1.5) * float32(0.0)) - 1.5) >= (3.14159 - (-2.25 + (3.14159 * 1.5)))) {
        v42 := GenS2{F0: 32767}
        v43 := GenS2{F0: 4294967295}
        _ = v42
        _ = v43
    } else {
        var v44 uint64 = (v22 * (v22 ^ v22))
        _ = v44
    }
    v45 := []int{2147483647, -9223372036854775808, 42}
    v46 := GenS3{F0: v3, F1: 54, F2: 249}
    v47 := GenS2{F0: -118}
    if false {
        switch "\n\t" {
        case "\n\t":
            switch "" {
            case "hello":
                var v48 int16 = 0
                _ = v48
            }
        default:
            v49 := GenS2{F0: 32768}
            _ = v49
        }
        v45 = append(v45, 1)
        var v50 int16 = -32768
        _ = v50
    }
    v22 += v22
    for fi51 := 0; fi51 < 7; fi51++ {
        if len(v45) > 0 {
            v45[0] = (fi51 << ((705) & 63))
        }
        _ = fi51
    }
    var v52 bool = (v1 == (v4 | v4))
    var v53 bool = false
    v52 = (!(!((v3 > v3) || false)))
    for fi54 := 0; fi54 < 5; fi54++ {
        for fi55 := 0; fi55 < 11; fi55++ {
            v22 = 1
            v56 := GenS3{F0: (v3 << ((377) & 63)), F1: 122, F2: (v3 | v3)}
            _ = fi55
            _ = v56
        }
        var v57 uint = genFn0(2147483647, -1, "a")
        v45 = append(v45, (-(fi54)))
        _ = fi54
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v52
    _ = v53
}
