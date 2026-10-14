package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int }
type GenS2 struct { F0 int8; F1 int16; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint32; F2 int8 }
type GenS5 struct { F0 int64; F1 int64; F2 uint64 }

func main() {
    var v1 uint = 2
    var v2 float64 = 0.0
    var v3 uint64 = 2
    v4 := []string{"\n\t", ("hello" + "\n\t"), "foo bar", "hello"}
    var v5 bool = true
    fmt.Println(v1)
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint32 = 340
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: (v5 && (!v5)), F1: 100, F2: fi6}
        switch "\n\t" {
        case "\n\t":
            var v10 uint8 = 2
            _ = v10
        case "":
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 uint32 = 690
                var v13 float32 = 0.1
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            fmt.Println(v1)
        }
        _ = fi6
        _ = v9
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 string = "foo bar"
        _ = len(v4)
        _ = fi14
        _ = v15
    }
    if v5 {
        fmt.Println(v2)
        v2 = v2
        var v16 uint = (v1 | (v1 + (v1 - 255)))
        _ = v16
    } else {
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint16 = 65535
            switch ("\n\t" + ("hello" + "")) {
            case "":
                if v5 {
                    fmt.Println(fi17, v5)
                    if true {
                        var v19 uint16 = v18
                        v4 = append(v4, "foo bar")
                        if true {
                            for fi20 := 0; fi20 < 10; fi20++ {
                                switch fi20 {
                                case 9223372036854775807:
                                    if (!(v5 || (!v5))) {
                                        v21 := GenS2{F0: -92, F1: 60, F2: (^(v1))}
                                        v22 := GenS1{F0: (42 > 42), F1: -1, F2: -386}
                                        _ = v21
                                        _ = v22
                                    } else {
                                        switch "a" {
                                        case "":
                                            v4 = append(v4, "hello")
                                        case "a":
                                            var v23 int32 = 2147483647
                                            _ = v23
                                        case "hello":
                                            v18 = (^(v18))
                                        default:
                                            fmt.Println(fi17, fi17, v2)
                                        }
                                        var v24 float64 = v2
                                        _ = v24
                                    }
                                case -911:
                                    v18 = 1
                                case 986:
                                    if len(v4) > 0 {
                                        v4[0] = ("a" + "a")
                                    }
                                default:
                                    var v25 uint = 0
                                    _ = v25
                                }
                                fmt.Println(v18, v18, fi20)
                                var v26 int = (^(fi17))
                                _ = fi20
                                _ = v26
                            }
                            if len(v4) > 0 {
                                v4[0] = ("foo bar" + "hello")
                            }
                            if true {
                                v2 -= (0.0 + v2)
                                for fi27 := 0; fi27 < 20; fi27++ {
                                    var v28 int8 = -79
                                    _ = fi27
                                    _ = v28
                                }
                                var v29 uint64 = (^(v3))
                                _ = v29
                            } else {
                                fmt.Println(v5)
                                if true {
                                    fmt.Println(v3, v18, v5)
                                } else {
                                    switch (fi17 >> ((v1) & 63)) {
                                    case -967:
                                        v2 = ((v2 * v2) + v2)
                                    }
                                }
                                v4 = append(v4, "a")
                            }
                        } else {
                            fmt.Println(fi17, v3, v3)
                            v30 := GenS2{F0: 42, F1: 575, F2: 18446744073709551615}
                            var v31 int16 = 32767
                            _ = v30
                            _ = v31
                        }
                        _ = v19
                    } else {
                        switch "" {
                        case "a":
                            var v32 int8 = int8(v18)
                            _ = v32
                        }
                        v33 := GenS2{F0: 0, F1: 694, F2: v1}
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        _ = v33
                    }
                    var v34 float32 = 3.14159
                    _ = v34
                }
            case "hello":
                v2 = v2
            case "a":
                v3 = v3
            }
            v1 += (v1 | (v1 | 255))
            _ = fi17
            _ = v18
        }
        v2 -= v2
    }
    var v35 int64 = -9223372036854775808
    fmt.Println(v3, v2, v5)
    var v36 int16 = 483
    if v5 {
        for fi37 := 0; fi37 < 1; fi37++ {
            v2 += (0.0 - v2)
            _ = fi37
        }
        v38 := GenS3{F0: 0}
        var v39 uint16 = 255
        _ = v38
        _ = v39
    }
    _ = len(v4)
    var v40 uint64 = (v3 + 4294967295)
    v36 += -1
    for fi41 := 0; fi41 < 18; fi41++ {
        var v42 float32 = 1e10
        _ = fi41
        _ = v42
    }
    v35 -= -9223372036854775808
    for fi43 := 0; fi43 < 1; fi43++ {
        if (!((fi43 ^ fi43) == 0)) {
            v1 += (v1 & (^(v1)))
        } else {
            if (v1 == 4294967295) {
                fmt.Println(v36, v35)
                v5 = v5
            } else {
                switch (fi43 >> (((v1 * (v1 - (v1 + (v1 | v1))))) & 63)) {
                case -2147483648:
                    v44 := GenS1{F0: (true && (false || v5)), F1: 127, F2: (fi43 * (fi43 >> ((473) & 63)))}
                    _ = v44
                case 1:
                    var v45 string = ("a" + ("\n\t" + "foo bar"))
                    _ = v45
                }
                v40 += 346
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v46 := GenS3{F0: 139}
            for fi47 := 0; fi47 < 9; fi47++ {
                switch "hello" {
                case "a":
                    v1 -= (v1 | 1)
                case "foo bar":
                    v4 = append(v4, "foo bar")
                case "":
                    if len(v4) > 0 {
                        v4[0] = (("foo bar" + "\n\t") + ("foo bar" + ""))
                    }
                    v35 -= 1
                }
                var v48 uint16 = 634
                var v49 float64 = 1e10
                v50 := GenS4{F0: -1, F1: 2, F2: -1}
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
            }
            _ = v46
        }
        v40 -= v3
        _ = fi43
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        v52 := GenS5{F0: int64(fi51), F1: 0, F2: 627}
        v1 = v1
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 4; fi53++ {
        v36 = v36
        _ = fi53
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
}
