package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 int16 = -1
    var v2 int64 = 9223372036854775807
    var v3 string = ("hello" + ("hello" + ""))
    var v4 uint8 = 2
    v3 = "hello"
    v5 := []int{9223372036854775807, -2147483648, 2147483647, 32767}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            if len(v5) > 0 {
                v5[0] = fi6
            }
            v2 -= 0
            _ = fi7
        }
        if true {
            if (!(!true)) {
                for fi8 := 0; fi8 < 11; fi8++ {
                    v3 = ((("" + ((v3 + v3) + (v3 + v3))) + ((v3 + "") + v3)) + "a")
                    _ = fi8
                }
            } else {
                if (false && true) {
                    v5 = append(v5, fi6)
                    var v9 float32 = 0.0
                    _ = v9
                }
                var v10 int = -662
                _ = v10
            }
            var v11 int16 = v1
            _ = v11
        } else {
            var v12 bool = true
            switch fi6 {
            case 9223372036854775807:
                if (v12 && (0 < 127)) {
                    v5 = append(v5, fi6)
                } else {
                    switch fi6 {
                    case 235:
                        v4 += (v4 + 0)
                    case -82:
                        var v13 uint16 = 827
                        _ = v13
                    case 253:
                        var v14 string = v3
                        _ = v14
                    default:
                        v15 := GenS1{F0: v1}
                        _ = v15
                    }
                    var v16 float32 = ((0.1 / 1.0) - -2.25)
                    _ = v16
                }
            case -9223372036854775808:
                v12 = (!true)
            }
            if true {
                for fi17 := 0; fi17 < 6; fi17++ {
                    switch v3 {
                    case "foo bar":
                        var v18 string = ("\n\t" + (v3 + v3))
                        _ = v18
                    }
                    var v19 uint = 65535
                    var v20 float64 = 3.14159
                    switch v3 {
                    case "a":
                        if v12 {
                            fmt.Println(v1, v4, v20)
                            if false {
                                if ((1e10 - (v20 - v20)) > 0.1) {
                                    v12 = v12
                                    var v21 uint = v19
                                    var v22 int = (fi17 | (fi6 | fi17))
                                    _ = v21
                                    _ = v22
                                } else {
                                    fmt.Println(fi17, fi17, v4)
                                    v5 = append(v5, (fi17 << ((2) & 63)))
                                }
                            } else {
                                _ = len(v5)
                                if ("" == v3) {
                                    var v23 uint = (v19 ^ v19)
                                    _ = v23
                                } else {
                                    v1 = (-(v1))
                                }
                                var v24 uint64 = 1
                                v2 = v2
                                v12 = (922 <= 224)
                                _ = v24
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v25 uint16 = 65535
                            _ = v25
                        }
                    case "":
                        var v26 int8 = -30
                        _ = v26
                    case "\n\t":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        switch (^(fi6)) {
                        case -1:
                            v2 = v2
                        default:
                            v20 += (v20 / v20)
                        }
                    }
                    fmt.Println(v1)
                    var v27 int32 = 2147483647
                    v20 -= 1e10
                    _ = fi17
                    _ = v19
                    _ = v20
                    _ = v27
                }
                for fi28 := 0; fi28 < 2; fi28++ {
                    v12 = (((v3 <= "") && v12) && v12)
                    if v12 {
                        v3 = (v3 + "foo bar")
                    } else {
                        fmt.Println(v4, v4, fi28)
                        v4 += 42
                    }
                    var v29 int = fi6
                    _ = fi28
                    _ = v29
                }
            } else {
                var v30 bool = v12
                v30 = (89 == 255)
                fmt.Println(v3, v30)
                _ = v30
            }
            _ = v12
        }
        var v31 uint8 = v4
        _ = fi6
        _ = v31
    }
    if true {
        var v32 float64 = -2.25
        v4 -= 1
        _ = v32
    }
    if (("foo bar" + "a") < (("" + ((v3 + v3) + ("" + (v3 + v3)))) + "a")) {
        if true {
            v3 = "a"
            if (1 > 65535) {
                v33 := GenS2{F0: -87}
                switch 128 {
                case 324:
                    for fi34 := 0; fi34 < 2; fi34++ {
                        var v35 uint = 255
                        v1 = v1
                        v2 = 42
                        _ = fi34
                        _ = v35
                    }
                case -2147483648:
                    if (v3 < "") {
                        if true {
                            switch v3 {
                            case "":
                                _ = len(v5)
                            case "hello":
                                var v36 uint8 = v4
                                _ = v36
                            case "\n\t":
                                v4 = 42
                            }
                            switch ((v3 + (v3 + v3)) + v3) {
                            case "foo bar":
                                if false {
                                    fmt.Println(v2)
                                    v37 := GenS1{F0: (v1 - 32767)}
                                    _ = v37
                                }
                            case "a":
                                var v38 bool = true
                                _ = v38
                            case "\n\t":
                                v39 := GenS3{F0: 1}
                                _ = v39
                            }
                            fmt.Println(v3, v2)
                        }
                    } else {
                        fmt.Println(v4, v3)
                        v3 = (("hello" + ("hello" + v3)) + v3)
                    }
                case 9223372036854775807:
                    switch (v3 + (v3 + (v3 + "hello"))) {
                    case "hello":
                        if true {
                            for fi40 := 0; fi40 < 16; fi40++ {
                                for fi41 := 0; fi41 < 7; fi41++ {
                                    var v42 uint8 = (v4 ^ 1)
                                    _ = fi41
                                    _ = v42
                                }
                                if true {
                                    var v43 uint = 4294967295
                                    _ = v43
                                } else {
                                    var v44 int = fi40
                                    v3 = ""
                                    fmt.Println(v2)
                                    _ = v44
                                }
                                var v45 string = (v3 + "foo bar")
                                _ = fi40
                                _ = v45
                            }
                        } else {
                            if (!(388 == 194)) {
                                for fi46 := 0; fi46 < 3; fi46++ {
                                    v2 = 42
                                    _ = fi46
                                }
                            } else {
                                for fi47 := 0; fi47 < 9; fi47++ {
                                    v3 = "hello"
                                    var v48 float64 = -2.25
                                    var v49 float32 = (1e10 - 1.5)
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            }
                        }
                    case "\n\t":
                        v2 -= v2
                    case "a":
                        fmt.Println(v4)
                    default:
                        v50 := GenS4{F0: 966}
                        _ = v50
                    }
                default:
                    v51 := GenS3{F0: 4294967295}
                    _ = v51
                }
                _ = v33
            } else {
                v3 = (v3 + (v3 + v3))
            }
        }
        v52 := GenS1{F0: (v1 & (v1 ^ v1))}
        _ = v52
    }
    var v53 uint32 = 1
    var v54 int32 = 282
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
}
