package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16) int16 {
    if true {
        if (255 <= 310) {
            var v4 int8 = 42
            _ = v4
        }
        fmt.Println(p0)
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 float64 = 1.5
            _ = fi5
            _ = v6
        }
    } else {
        fmt.Println(p0)
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                p0 -= (p0 >> ((2) & 63))
                p0 -= (p0 >> ((255) & 63))
                var v9 int32 = 0
                _ = fi8
                _ = v9
            }
            switch (fi7 ^ fi7) {
            case 0:
                switch "hello" {
                case "a":
                    switch (^(fi7)) {
                    case 853:
                        if false {
                            switch "hello" {
                            case "":
                                var v10 int = fi7
                                _ = v10
                            case "hello":
                                p0 -= -32768
                            case "\n\t":
                                v11 := []int{(fi7 - (^(fi7)))}
                                _ = v11
                            default:
                                p0 = p0
                            }
                            p0 += 724
                            v12 := GenS1{F0: 32767}
                            _ = v12
                        } else {
                            if (("a" + ("" + ("\n\t" + ("" + ("foo bar" + "hello"))))) == "hello") {
                                for fi13 := 0; fi13 < 12; fi13++ {
                                    var v14 int32 = 878
                                    var v15 uint = 18446744073709551615
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            }
                        }
                    case -598:
                        v16 := GenS1{F0: 1}
                        _ = v16
                    }
                case "foo bar":
                    if false {
                        p0 = p0
                        switch 0 {
                        case -4:
                            var v17 int64 = 554
                            _ = v17
                        case 0:
                            v18 := GenS1{F0: 918}
                            _ = v18
                        case 956:
                            fmt.Println(fi7, p0)
                        default:
                            for fi19 := 0; fi19 < 13; fi19++ {
                                v20 := GenS1{F0: 32767}
                                _ = fi19
                                _ = v20
                            }
                        }
                        v21 := []int{(fi7 << ((797) & 63)), 32768, (fi7 << ((1) & 63))}
                        _ = v21
                    }
                case "\n\t":
                    v22 := []int{fi7, (fi7 + fi7), 4294967295, (fi7 * 1)}
                    _ = v22
                default:
                    v23 := []int{-9223372036854775808, -1}
                    _ = v23
                }
            case 340:
                p0 = (p0 << ((520) & 63))
            case -212:
                p0 = p0
            }
            p0 += p0
            _ = fi7
        }
    }
    v24 := []int{4294967295, 128, 42}
    _ = p0
    _ = v24
    return (p0 >> ((134) & 63))
}

func main() {
    var v1 string = "hello"
    var v2 uint32 = 42
    v3 := make(map[string]bool)
    for fi25 := 0; fi25 < 4; fi25++ {
        delete(v3, v1)
        fmt.Println(v2, v1, v2)
        _ = fi25
    }
    v1 = (((v1 + v1) + (v1 + "\n\t")) + v1)
    var v26 int8 = -1
    for fi27 := 0; fi27 < 3; fi27++ {
        if true {
            switch v1 {
            case "":
                v28 := GenS1{F0: 355}
                _ = v28
            case "\n\t":
                v29 := GenS1{F0: 2}
                _ = v29
            case "a":
                v1 = v1
            default:
                if false {
                    switch 0 {
                    case 9223372036854775807:
                        v2 = (v2 - 678)
                    case 4294967295:
                        fmt.Println(v1)
                    case -9223372036854775808:
                        switch "a" {
                        case "a":
                            for fi30 := 0; fi30 < 2; fi30++ {
                                var v31 int32 = 42
                                v31 += -413
                                _ = fi30
                                _ = v31
                            }
                        case "\n\t":
                            switch 1 {
                            case -9223372036854775808:
                                fmt.Println(v26)
                            case 627:
                                v26 -= v26
                            default:
                                var v32 uint64 = 18446744073709551615
                                _ = v32
                            }
                        }
                    }
                    var v33 int32 = 2147483647
                    _ = v33
                }
            }
            for fi34 := 0; fi34 < 6; fi34++ {
                var v35 int8 = v26
                _ = fi34
                _ = v35
            }
        }
        v36 := GenS1{F0: 554}
        _ = fi27
        _ = v36
    }
    if (42 <= 839) {
        for fi37 := 0; fi37 < 5; fi37++ {
            fmt.Println(fi37)
            _ = fi37
        }
        v1 = v1
        v38 := GenS2{F0: 239}
        _ = v38
    } else {
        v39 := GenS3{F0: 42}
        var v40 uint64 = 687
        _ = v39
        _ = v40
    }
    if false {
        var v41 float32 = (float32(-2.25) + ((1e10 + 0.0) * (1e10 / 1e10)))
        _ = v41
    }
    for fi42 := 0; fi42 < 5; fi42++ {
        var v43 int8 = (^(v26))
        _ = fi42
        _ = v43
    }
    v44 := v3[v1]
    var v45 int16 = 32767
    var v46 int = -961
    var v47 uint16 = 255
    fmt.Println(v2, v45)
    fmt.Println(v1)
    fmt.Println(v26)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
