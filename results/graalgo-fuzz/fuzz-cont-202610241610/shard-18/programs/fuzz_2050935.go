package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int, p1 uint32, p2 int16) bool {
    var v7 int64 = -9223372036854775808
    if true {
        v8 := []int{}
        var v9 bool = (586 > 255)
        _ = v8
        _ = v9
    } else {
        if false {
            for fi10 := 0; fi10 < 2; fi10++ {
                v7 = (-(v7))
                _ = fi10
            }
            v11 := []int{}
            _ = v11
        } else {
            var v12 string = ""
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}
func genFn1() uint32 {
    v13 := GenS1{F0: 845, F1: "hello"}
    var v14 uint = 4294967295
    v14 -= v14
    switch -714 {
    case 32768:
        v14 = v14
    case 128:
        v14 = 255
    default:
        v15 := GenS1{F0: 187, F1: "hello"}
        _ = v15
    }
    _ = v13
    _ = v14
    return 147
}

func main() {
    var v1 bool = true
    var v2 uint16 = 383
    var v3 int32 = 0
    var v4 uint64 = 2
    var v5 uint16 = (^(v2))
    var v6 uint16 = v5
    v16 := []int{int(v5), 4294967295, -1}
    var v17 int16 = -1
    var v18 int = -9223372036854775808
    v19 := GenS2{F0: (!true), F1: 0}
    var v20 int32 = v3
    if v1 {
        switch (-(v18)) {
        case 9223372036854775807:
            v6 -= v2
        case 128:
            var v21 int64 = 1
            _ = v21
        case -2147483648:
            var v22 uint16 = (v6 >> ((59) & 63))
            _ = v22
        }
        _ = len(v16)
        if false {
            var v23 float32 = (float32(1e10) / -2.25)
            _ = v23
        } else {
            if (2147483647 > v3) {
                if true {
                    var v24 float32 = -2.25
                    v25 := GenS1{F0: 0, F1: ""}
                    if (!(v1 || (1 < v5))) {
                        var v26 bool = false
                        if (v1 && ((1e10 >= ((0.1 + (0.1 / (1.5 / 1.0))) + 3.14159)) && (!(2147483648 != 652)))) {
                            fmt.Println(v5)
                            if v26 {
                                v6 -= (v5 + 668)
                            }
                        }
                        for fi27 := 0; fi27 < 19; fi27++ {
                            switch int(v20) {
                            case -954:
                                fmt.Println(fi27, v6, v4)
                            default:
                                var v28 float64 = 0.1
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 18; fi29++ {
                                switch int(v5) {
                                case 9223372036854775807:
                                    var v30 float32 = 0.1
                                    _ = v30
                                case 2147483647:
                                    v6 += 0
                                default:
                                    fmt.Println(v1, v6)
                                }
                                if len(v16) > 0 {
                                    v16[0] = v18
                                }
                                _ = fi29
                            }
                            if false {
                                var v31 string = "foo bar"
                                _ = v31
                            }
                            _ = fi27
                        }
                        _ = v26
                    } else {
                        var v32 int16 = v17
                        for fi33 := 0; fi33 < 7; fi33++ {
                            switch 32767 {
                            case -651:
                                switch (((("foo bar" + ("" + "hello")) + ("a" + "foo bar")) + ("hello" + "foo bar")) + (("\n\t" + "") + "")) {
                                case "":
                                    var v34 int = v18
                                    _ = v34
                                default:
                                    v4 += (v4 ^ 42)
                                }
                            case 209:
                                v32 += (^(v32))
                            case -219:
                                v20 = v3
                            default:
                                v35 := GenS1{F0: 782, F1: "\n\t"}
                                _ = v35
                            }
                            _ = fi33
                        }
                        _ = v32
                    }
                    _ = v24
                    _ = v25
                } else {
                    v36 := GenS1{F0: genFn1(), F1: ""}
                    var v37 uint = 448
                    for fi38 := 0; fi38 < 5; fi38++ {
                        fmt.Println(v4, fi38, fi38)
                        _ = fi38
                    }
                    _ = v36
                    _ = v37
                }
            } else {
                fmt.Println(v20)
                if len(v16) > 0 {
                    v16[0] = (v18 - v18)
                }
                for fi39 := 0; fi39 < 12; fi39++ {
                    v40 := GenS1{F0: 389, F1: "a"}
                    _ = fi39
                    _ = v40
                }
            }
        }
    } else {
        fmt.Println(v4)
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        if false {
            switch v18 {
            case 4294967295:
                var v41 int64 = 9223372036854775807
                _ = v41
            case 32767:
                v42 := GenS1{F0: 42, F1: "hello"}
                _ = v42
            case 1:
                _ = len(v16)
            }
        } else {
            fmt.Println(v18, v3)
            if v1 {
                v4 = v4
                v43 := GenS3{F0: 94}
                _ = v43
            }
            var v44 uint16 = 418
            _ = v44
        }
    }
    var v45 int16 = (v17 | (v17 - (-(v17))))
    v46 := GenS3{F0: 2}
    v47 := GenS1{F0: 2, F1: ""}
    fmt.Println(v6, v3)
    var v48 bool = true
    for fi49 := 0; fi49 < 11; fi49++ {
        v50 := GenS1{F0: 2147483647, F1: "a"}
        var v51 uint = 0
        v51 = v51
        _ = fi49
        _ = v50
        _ = v51
    }
    v52 := GenS3{F0: 128}
    v5 -= v2
    if (!(523 < uint32(v4))) {
        switch (v18 & -1) {
        case -660:
            var v53 int32 = v3
            _ = v53
        }
        var v54 int = (v18 % (((v18 - v18)) | 1))
        if false {
            var v55 uint = 1
            v54 = v54
            switch ("hello" + ("a" + (("foo bar" + ("a" + "hello")) + (("" + ((((("\n\t" + "\n\t") + ("\n\t" + "")) + "\n\t") + "\n\t") + ("foo bar" + (("a" + "hello") + "foo bar")))) + ("a" + "foo bar"))))) {
            case "foo bar":
                var v56 float64 = 0.1
                _ = v56
            default:
                var v57 int8 = 1
                _ = v57
            }
            _ = v55
        } else {
            v58 := GenS1{F0: 208, F1: ("foo bar" + "")}
            _ = v58
        }
        _ = v54
    }
    for fi59 := 0; fi59 < 8; fi59++ {
        v3 = (v3 - 1)
        for fi60 := 0; fi60 < 15; fi60++ {
            fmt.Println(v5, v17)
            _ = fi60
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v52
    _ = v6
}
