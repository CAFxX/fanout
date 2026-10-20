package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := []uint32{567, 2}
    v2 := []string{"a", "hello", "\n\t", ("" + "foo bar")}
    var v3 uint = 65535
    var v4 int = -465
    var v5 int8 = 42
    var v6 int = v4
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: ""}
        var v9 int64 = 9223372036854775807
        var v10 int16 = -32768
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v2 = append(v2, "\n\t")
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint64 = 2
        var v13 int32 = -2147483648
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int8 = v5
    switch (^(v4)) {
    case -9223372036854775808:
        v15 := GenS1{F0: "a"}
        _ = v15
    case -920:
        if true {
            if false {
                v16 := GenS1{F0: ("\n\t" + "foo bar")}
                _ = v16
            } else {
                if len(v2) > 0 {
                    v2[0] = ("a" + "")
                }
                if len(v1) > 0 {
                    v1[0] = 0
                }
                if ("a" != "a") {
                    fmt.Println(v3)
                    v5 += 42
                } else {
                    var v17 uint = 1
                    switch ("a" + "\n\t") {
                    case "hello":
                        switch 161 {
                        case -2147483648:
                            if false {
                                var v18 uint32 = 65535
                                for fi19 := 0; fi19 < 4; fi19++ {
                                    for fi20 := 0; fi20 < 16; fi20++ {
                                        var v21 uint16 = 255
                                        v18 = (v18 + (v18 / ((v18) | 1)))
                                        fmt.Println(fi19, v6)
                                        _ = fi20
                                        _ = v21
                                    }
                                    _ = fi19
                                }
                                _ = v18
                            }
                        default:
                            for fi22 := 0; fi22 < 11; fi22++ {
                                v17 -= v3
                                _ = fi22
                            }
                        }
                    case "":
                        v23 := GenS1{F0: "foo bar"}
                        _ = v23
                    case "a":
                        fmt.Println(v6, v14)
                    }
                    for fi24 := 0; fi24 < 19; fi24++ {
                        for fi25 := 0; fi25 < 19; fi25++ {
                            v6 = 0
                            _ = fi25
                        }
                        switch v4 {
                        case 2147483647:
                            var v26 uint16 = 0
                            _ = v26
                        case 148:
                            v27 := GenS1{F0: "hello"}
                            _ = v27
                        case 32767:
                            fmt.Println(v17, v17)
                        default:
                            v4 += -893
                        }
                        if true {
                            for fi28 := 0; fi28 < 7; fi28++ {
                                v29 := GenS1{F0: ("" + ("\n\t" + "foo bar"))}
                                v2 = append(v2, "a")
                                _ = fi28
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 5; fi30++ {
                                v3 += 0
                                v31 := GenS1{F0: (("hello" + ((("\n\t" + ("" + "\n\t")) + "a") + ("" + ""))) + ("a" + "\n\t"))}
                                switch v4 {
                                case -738:
                                    if true {
                                        fmt.Println(v4, v6, v3)
                                        v3 -= 1
                                    } else {
                                        var v32 float64 = -2.25
                                        _ = v32
                                    }
                                case -2147483648:
                                    v6 -= int(v5)
                                case 127:
                                    switch "a" {
                                    case "hello":
                                        switch 0 {
                                        case 2147483647:
                                            var v33 uint64 = 18446744073709551615
                                            _ = v33
                                        case 42:
                                            v34 := GenS1{F0: "\n\t"}
                                            _ = v34
                                        case -2147483648:
                                            if true {
                                                fmt.Println(fi24, v3, v14)
                                                switch v6 {
                                                case -1:
                                                    v35 := GenS1{F0: "foo bar"}
                                                    var v36 float32 = 0.0
                                                    var v37 int16 = 42
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                case 32767:
                                                    var v38 uint8 = 255
                                                    _ = v38
                                                }
                                            }
                                            v14 = (v5 * (-(v5)))
                                        default:
                                            v5 -= 127
                                        }
                                    case "a":
                                        switch "foo bar" {
                                        case "foo bar":
                                            v39 := GenS2{F0: -1, F1: 0}
                                            v40 := GenS3{F0: 332}
                                            _ = v39
                                            _ = v40
                                        }
                                    case "foo bar":
                                        fmt.Println(v14)
                                    default:
                                        v41 := GenS1{F0: ""}
                                        _ = v41
                                    }
                                default:
                                    var v42 string = "\n\t"
                                    _ = v42
                                }
                                _ = fi30
                                _ = v31
                            }
                        }
                        _ = fi24
                    }
                    _ = v17
                }
            }
            v43 := GenS2{F0: -1, F1: 66}
            _ = v43
        } else {
            v6 = v6
            v44 := GenS3{F0: (v6 << ((v3) & 63))}
            fmt.Println(v3)
            _ = v44
        }
    case 32767:
        var v45 int32 = -2147483648
        _ = v45
    default:
        _ = len(v2)
    }
    v46 := GenS3{F0: (v4 + -9223372036854775808)}
    if true {
        switch (v6 | (v4 >> ((v3) & 63))) {
        case 32768:
            v47 := GenS2{F0: -384, F1: 127}
            _ = v47
        }
    }
    var v48 int = (v4 - v4)
    v48 += v6
    v1 = append(v1, 65535)
    var v49 int32 = 392
    if (!true) {
        var v50 uint = (v3 & (v3 ^ (^(v3))))
        for fi51 := 0; fi51 < 5; fi51++ {
            fmt.Println(v5, v6)
            if true {
                for fi52 := 0; fi52 < 18; fi52++ {
                    if (!false) {
                        v53 := GenS4{F0: 4294967295}
                        _ = v53
                    }
                    v54 := GenS5{F0: int64(v3)}
                    _ = fi52
                    _ = v54
                }
                v3 += v50
                if true {
                    for fi55 := 0; fi55 < 6; fi55++ {
                        fmt.Println(v14)
                        if (2 == 2) {
                            v49 -= (v49 >> ((v3) & 63))
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi55
                    }
                }
            } else {
                if false {
                    v2 = append(v2, "hello")
                    if (426 >= 18446744073709551615) {
                        for fi56 := 0; fi56 < 16; fi56++ {
                            fmt.Println(v49)
                            _ = fi56
                        }
                    }
                }
                v5 += (^(v5))
            }
            _ = fi51
        }
        _ = v50
    } else {
        v4 -= v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
