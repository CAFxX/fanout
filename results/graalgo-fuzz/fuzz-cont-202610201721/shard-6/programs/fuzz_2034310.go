package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int16; F1 int; F2 uint }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 int16 = 69
    var v2 int16 = (-(v1))
    var v3 int = -413
    v4 := []string{""}
    var v5 bool = true
    v6 := make(map[string]string)
    if v5 {
        v1 += (-(v2))
        var v7 string = ""
        v8 := GenS1{F0: 4294967295}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 12; fi9++ {
            for fi10 := 0; fi10 < 2; fi10++ {
                fmt.Println(v2, fi9, v2)
                delete(v6, "hello")
                v6[""] = ("hello" + ("\n\t" + "foo bar"))
                _ = fi10
            }
            switch fi9 {
            case 548:
                v3 = fi9
            }
            v2 = -1
            _ = fi9
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            if false {
                if v5 {
                    for fi12 := 0; fi12 < 17; fi12++ {
                        var v13 uint32 = 42
                        fmt.Println(v13, fi11)
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 10; fi14++ {
                        fmt.Println(v5, v3)
                        fmt.Println(v1, v5)
                        _ = len(v6)
                        _ = fi14
                    }
                    var v15 int32 = 4
                    _ = v15
                }
                v16 := GenS1{F0: 2}
                _ = v16
            }
            var v17 string = ((("\n\t" + "\n\t") + "\n\t") + "")
            _ = fi11
            _ = v17
        }
    }
    var v18 uint64 = 1
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 uint64 = v18
        fmt.Println(v20, v3, v1)
        _ = fi19
        _ = v20
    }
    delete(v6, "a")
    var v21 float64 = float64(1.5)
    if false {
        v4 = append(v4, "hello")
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v2 += -32768
        _ = fi22
    }
    if ((128 < 1) && false) {
        if (1.5 != 0.1) {
            if (((("hello" + "foo bar") + "foo bar") > "hello") || true) {
                v2 = 152
                fmt.Println(v21)
                fmt.Println(v3, v1)
            }
        }
        fmt.Println(v5)
        v3 += v3
    }
    if (9223372036854775807 < -701) {
        v23 := GenS2{F0: v3}
        v21 = v21
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 uint8 = 42
            _ = len(v6)
            if (!false) {
                var v26 uint8 = v25
                switch "\n\t" {
                case "\n\t":
                    var v27 int16 = v1
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 10; fi28++ {
                        switch -276 {
                        case 127:
                            var v29 int8 = -69
                            _ = v29
                        }
                        fmt.Println(v3, v1)
                        var v30 uint64 = uint64(v25)
                        fmt.Println(v18)
                        _ = fi28
                        _ = v30
                    }
                }
                v1 += int16(v18)
                _ = v26
            } else {
                if v5 {
                    v31 := GenS3{F0: int(v25), F1: (v2 - -32768)}
                    _ = v31
                }
                switch -1 {
                case 32768:
                    _ = len(v6)
                case -570:
                    var v32 int16 = v2
                    _ = v32
                default:
                    if (v5 || v5) {
                        for fi33 := 0; fi33 < 9; fi33++ {
                            v25 += 128
                            if len(v4) > 0 {
                                v4[0] = ("foo bar" + ("foo bar" + "\n\t"))
                            }
                            if (v25 < v25) {
                                if (true && false) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    fmt.Println(v1, v21)
                                    switch -1 {
                                    case 2147483647:
                                        var v34 int64 = -9223372036854775808
                                        _ = v34
                                    case 128:
                                        v35 := v6[("foo bar" + "hello")]
                                        _ = v35
                                    case 9223372036854775807:
                                        if ((v2 - 618) == -506) {
                                            if false {
                                                v4 = append(v4, "hello")
                                                fmt.Println(v3)
                                            }
                                            v5 = (0.0 <= (((v21 - v21) - (v21 + ((float64((((3.14159 + 1e10) - -2.25) * 0.1)) - (v21 - 3.14159)) * 1e10))) - v21))
                                            v36 := GenS4{F0: v2, F1: v3, F2: uint(v25)}
                                            _ = v36
                                        }
                                    }
                                    v5 = ((v5 || (858 <= 255)) && true)
                                } else {
                                    if v5 {
                                        v2 -= 32767
                                    }
                                    if (!(4294967295 == 4294967295)) {
                                        if v5 {
                                            var v37 int = fi24
                                            fmt.Println(v1, v21)
                                            if v5 {
                                                var v38 uint32 = 2147483648
                                                v39 := GenS3{F0: 1, F1: -32768}
                                                fmt.Println(v3)
                                                _ = v38
                                                _ = v39
                                            }
                                            _ = v37
                                        } else {
                                            _ = len(v4)
                                            switch "a" {
                                            case "foo bar":
                                                var v40 uint64 = v18
                                                var v41 uint16 = 255
                                                var v42 uint16 = v41
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            case "":
                                                delete(v6, ("a" + ("foo bar" + (("\n\t" + "a") + "hello"))))
                                            }
                                            switch "" {
                                            case "":
                                                if v5 {
                                                    v21 = v21
                                                } else {
                                                    var v43 float64 = v21
                                                    _ = v43
                                                }
                                            case "\n\t":
                                                v44 := GenS1{F0: 1}
                                                _ = v44
                                            }
                                        }
                                        if (-1 != -611) {
                                            v3 = (fi33 >> ((895) & 63))
                                            switch (fi24 - fi24) {
                                            case 32768:
                                                switch 128 {
                                                case 4294967295:
                                                    _ = len(v4)
                                                case 42:
                                                    v45 := GenS5{F0: 309}
                                                    _ = v45
                                                }
                                            case 128:
                                                v46 := v6["\n\t"]
                                                _ = v46
                                            }
                                            switch ("" + ("foo bar" + ("" + "hello"))) {
                                            case "a":
                                                var v47 int64 = 9223372036854775807
                                                _ = v47
                                            case "":
                                                fmt.Println(v18)
                                            case "foo bar":
                                                var v48 uint16 = 433
                                                _ = v48
                                            }
                                        } else {
                                            v25 = (^(v25))
                                            v5 = v5
                                            if true {
                                                v2 = (v1 / ((v1) | 1))
                                                v49 := GenS1{F0: 2147483648}
                                                _ = v49
                                            }
                                        }
                                    }
                                    v3 -= 32768
                                }
                                fmt.Println(v25, fi24)
                            }
                            _ = fi33
                        }
                        for fi50 := 0; fi50 < 16; fi50++ {
                            if true {
                                v18 += v18
                            }
                            if (!(-1 <= int8(v1))) {
                                if ("\n\t" == (("a" + ("a" + "\n\t")) + "hello")) {
                                    if (int32(v18) == 0) {
                                        var v51 float64 = 0.1
                                        switch fi50 {
                                        case 4294967295:
                                            var v52 string = ("" + ("" + ""))
                                            _ = v52
                                        default:
                                            var v53 int = (fi24 & (fi24 % ((fi24) | 1)))
                                            _ = v53
                                        }
                                        _ = v51
                                    } else {
                                        var v54 uint32 = 0
                                        v18 = v18
                                        _ = v54
                                    }
                                    fmt.Println(fi24)
                                }
                                v5 = true
                            }
                            _ = fi50
                        }
                    } else {
                        var v55 int8 = int8(v2)
                        switch (("hello" + "") + "") {
                        case "\n\t":
                            fmt.Println(fi24, fi24)
                        default:
                            delete(v6, ("foo bar" + "foo bar"))
                        }
                        _ = v55
                    }
                }
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    fmt.Println(v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
