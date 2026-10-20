package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (42 != 32768)
    v3 := make(map[int]int)
    var v4 int16 = 32767
    v5 := make(map[int]int)
    var v6 string = ""
    for fi7 := 0; fi7 < 14; fi7++ {
        switch v6 {
        case "hello":
            v4 -= 42
        default:
            fmt.Println(v6, v2, v2)
        }
        v2 = (!(4294967295 <= 4294967295))
        _ = fi7
    }
    if (v2 && false) {
        v4 -= (^(v4))
        var v8 uint16 = 65535
        v9 := v3[1]
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 5; fi11++ {
                v4 += v4
                _ = fi11
            }
            for fi12 := 0; fi12 < 4; fi12++ {
                if (704 >= 570) {
                    v2 = false
                } else {
                    v6 = (v6 + "foo bar")
                    v13 := GenS1{F0: 155, F1: 1, F2: 471}
                    _ = v13
                }
                v6 = v6
                v4 = (v4 - (^(v4)))
                _ = fi12
            }
            var v14 int = 32767
            _ = fi10
            _ = v14
        }
    }
    var v15 float32 = ((float32(-2.25) + 3.14159) - -2.25)
    v6 = "a"
    fmt.Println(v4, v6, v15)
    for fi16 := 0; fi16 < 9; fi16++ {
        if (127 < 13) {
            if (18446744073709551615 < 18446744073709551615) {
                if false {
                    _ = len(v3)
                }
            }
            switch (-(fi16)) {
            case 1:
                v17 := GenS2{F0: v2}
                _ = v17
            case 32767:
                v18 := GenS3{F0: 42}
                _ = v18
            case -597:
                switch fi16 {
                case 9223372036854775807:
                    for fi19 := 0; fi19 < 9; fi19++ {
                        var v20 int64 = 9223372036854775807
                        _ = fi19
                        _ = v20
                    }
                case -9223372036854775808:
                    v21 := GenS1{F0: -9223372036854775808, F1: 0, F2: 42}
                    _ = v21
                }
            default:
                switch fi16 {
                case 999:
                    for fi22 := 0; fi22 < 3; fi22++ {
                        if (((65535 == 551) || (447 < -887)) && true) {
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                switch (fi22 << ((255) & 63)) {
                                case -347:
                                    var v23 uint8 = 2
                                    _ = v23
                                case -358:
                                    if true {
                                        fmt.Println(v2, v15)
                                        v15 -= (v15 - v15)
                                        v6 = v6
                                    } else {
                                        fmt.Println(v6)
                                    }
                                case -1:
                                    var v24 int32 = 1
                                    _ = v24
                                }
                            case "a":
                                switch "a" {
                                case "a":
                                    var v25 int32 = 0
                                    _ = v25
                                case "":
                                    v4 = v4
                                default:
                                    if false {
                                        var v26 uint32 = 65535
                                        _ = v26
                                    } else {
                                        v27 := GenS3{F0: 42}
                                        _ = v27
                                    }
                                }
                            }
                            var v28 int64 = 0
                            if false {
                                v29 := GenS3{F0: 42}
                                _ = v29
                            }
                            v30 := []int{}
                            _ = v28
                            _ = v30
                        }
                        v31 := GenS2{F0: (v2 || true)}
                        var v32 uint8 = 0
                        _ = fi22
                        _ = v31
                        _ = v32
                    }
                case -98:
                    v2 = (202 == 4294967295)
                case -456:
                    var v33 int64 = int64(v4)
                    _ = v33
                default:
                    v15 += float32(0.1)
                }
            }
        } else {
            v15 -= 0.1
            var v34 uint32 = 2147483648
            _ = v34
        }
        _ = fi16
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        v36 := GenS4{F0: (v4 << ((18446744073709551615) & 63)), F1: 75}
        v37 := []int{(fi35 ^ 0), (fi35 << ((4294967295) & 63))}
        if (42 >= -2147483648) {
            v38 := v1[875]
            _ = v38
        } else {
            for fi39 := 0; fi39 < 18; fi39++ {
                if len(v37) > 0 {
                    v37[0] = fi39
                }
                switch ((("" + v6) + ((v6 + (v6 + v6)) + ((v6 + v6) + "\n\t"))) + v6) {
                case "\n\t":
                    var v40 int8 = 100
                    _ = v40
                case "foo bar":
                    if len(v37) >= 2 {
                        v37 = v37[1:len(v37)]
                    }
                case "a":
                    _ = len(v37)
                default:
                    if len(v37) > 0 {
                        v37[0] = (fi39 & fi35)
                    }
                }
                switch (fi39 << ((18446744073709551615) & 63)) {
                case 1:
                    v6 = v6
                case -529:
                    var v41 int32 = -252
                    _ = v41
                case -976:
                    v37 = append(v37, (fi35 << ((255) & 63)))
                }
                _ = fi39
            }
            if v2 {
                v37 = append(v37, 1)
                v2 = ((!false) || (127 == 255))
                var v42 uint16 = 65535
                _ = v42
            } else {
                for fi43 := 0; fi43 < 11; fi43++ {
                    v44 := GenS1{F0: 1, F1: 42, F2: 1}
                    fmt.Println(v4, v6)
                    v4 = (v4 << ((1) & 63))
                    _ = fi43
                    _ = v44
                }
                v2 = (!(("foo bar" + "hello") < (v6 + (v6 + v6))))
            }
            var v45 float64 = 1.5
            _ = v45
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    if (!v2) {
        if v2 {
            if (!(false || v2)) {
                fmt.Println(v15)
                for fi46 := 0; fi46 < 13; fi46++ {
                    v47 := GenS1{F0: 0, F1: -1, F2: 42}
                    _ = fi46
                    _ = v47
                }
                v6 = "\n\t"
            }
        } else {
            delete(v3, 160)
            var v48 int32 = 0
            if false {
                if v2 {
                    switch ((v6 + v6) + v6) {
                    case "a":
                        v2 = true
                    case "foo bar":
                        if (327 < 9223372036854775807) {
                            v15 -= 0.0
                        }
                    case "":
                        var v49 int64 = -9223372036854775808
                        _ = v49
                    }
                    if v2 {
                        if ((0 == 4294967295) || (!false)) {
                            if v2 {
                                fmt.Println(v4, v2, v6)
                            } else {
                                var v50 int64 = 770
                                _ = v50
                            }
                            v48 = 884
                            fmt.Println(v48, v48, v48)
                        }
                        v51 := GenS3{F0: 2}
                        _ = v51
                    }
                    v52 := GenS5{F0: "hello"}
                    _ = v52
                } else {
                    var v53 int64 = 315
                    v48 -= v48
                    _ = v53
                }
                var v54 float64 = 0.1
                _ = v54
            } else {
                var v55 int8 = 0
                if true {
                    var v56 int8 = (^(v55))
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 1; fi57++ {
                        if true {
                            v55 = -1
                        } else {
                            fmt.Println(v6)
                        }
                        _ = fi57
                    }
                    var v58 int = -2147483648
                    var v59 float64 = 3.14159
                    _ = v58
                    _ = v59
                }
                v60 := GenS4{F0: v4, F1: 127}
                _ = v55
                _ = v60
            }
            _ = v48
        }
        fmt.Println(v6, v4, v2)
        fmt.Println(v4, v6)
    }
    fmt.Println(v2)
    v3[550] = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
