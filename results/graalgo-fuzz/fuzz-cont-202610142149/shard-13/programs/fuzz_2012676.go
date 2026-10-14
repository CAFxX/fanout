package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int64; F2 bool }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16 }
func genFn0() uint {
    for fi7 := 0; fi7 < 18; fi7++ {
        switch ((((("foo bar" + ("\n\t" + "foo bar")) + "a") + "\n\t") + ("hello" + "hello")) + "hello") {
        case "hello":
            if (894 >= 0) {
                var v8 string = (("\n\t" + "foo bar") + "")
                _ = v8
            }
        case "foo bar":
            fmt.Println(fi7)
        }
        _ = fi7
    }
    switch "a" {
    case "foo bar":
        var v9 uint8 = 42
        _ = v9
    case "a":
        var v10 uint64 = 1
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint = 2
        var v13 bool = true
        var v14 uint32 = 1
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 int64 = 1
    _ = v15
    return 2
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[int]string)
    var v3 uint = 255
    var v4 uint = v3
    v5 := []uint32{184}
    v6 := []uint32{}
    v3 = 1
    if true {
        v4 = 903
    } else {
        v2[-166] = ("a" + "foo bar")
        var v16 uint16 = 32767
        _ = v16
    }
    var v17 int32 = 1
    if false {
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(v3, fi18, v17)
            v17 = 42
            _ = fi18
        }
        if (!(-9223372036854775808 != -988)) {
            fmt.Println(v4)
            var v19 uint8 = 0
            v19 = (^(v19))
            _ = v19
        }
        v20 := GenS1{F0: 558}
        _ = v20
    }
    if (!(!true)) {
        v4 -= 255
        for fi21 := 0; fi21 < 2; fi21++ {
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 uint64 = 576
                _ = fi22
                _ = v23
            }
            var v24 int = (fi21 & 1)
            var v25 uint8 = 127
            _ = fi21
            _ = v24
            _ = v25
        }
    } else {
        if (genFn0() > genFn0()) {
            if false {
                var v26 float32 = 1e10
                switch 360 {
                case 128:
                    var v27 int8 = 1
                    _ = v27
                case 2147483647:
                    var v28 int16 = 1
                    _ = v28
                case -1:
                    _ = len(v5)
                }
                v3 = (^(v4))
                _ = v26
            } else {
                v3 -= v4
            }
        } else {
            if true {
                fmt.Println(v3)
                v17 = (v17 & v17)
            } else {
                var v29 int64 = 42
                if false {
                    switch "\n\t" {
                    case "a":
                        var v30 float64 = -2.25
                        _ = v30
                    case "":
                        for fi31 := 0; fi31 < 1; fi31++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v32 := GenS2{F0: "", F1: -9223372036854775808, F2: ("" >= ("\n\t" + "\n\t"))}
                            _ = fi31
                            _ = v32
                        }
                    case "\n\t":
                        var v33 int32 = (v17 >> ((v4) & 63))
                        _ = v33
                    default:
                        switch 128 {
                        case -106:
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                        case -449:
                            fmt.Println(v3, v17)
                        }
                    }
                    fmt.Println(v4, v17, v3)
                    fmt.Println(v3, v3)
                } else {
                    v34 := v2[int(v4)]
                    var v35 float64 = (0.1 - 1e10)
                    _ = v34
                    _ = v35
                }
                var v36 string = ("a" + "foo bar")
                _ = v29
                _ = v36
            }
            v4 -= 65535
            v37 := GenS3{F0: 128, F1: true, F2: ("a" + (("\n\t" + ("" + "\n\t")) + "\n\t"))}
            _ = v37
        }
    }
    var v38 uint16 = 255
    for fi39 := 0; fi39 < 15; fi39++ {
        for fi40 := 0; fi40 < 8; fi40++ {
            if false {
                if true {
                    v41 := GenS4{F0: 569, F1: (("hello" + ("foo bar" + "a")) + ("foo bar" + ((((("a" + "hello") + ("a" + (("foo bar" + ("\n\t" + ("foo bar" + "a"))) + ((("" + "foo bar") + "") + "a")))) + "") + "") + "\n\t")))}
                    _ = v41
                } else {
                    var v42 uint8 = 29
                    var v43 uint64 = 42
                    switch (fi39 * (^(fi39))) {
                    case -400:
                        v2[int(v17)] = "a"
                    }
                    _ = v42
                    _ = v43
                }
                v17 = (^(v17))
                var v44 uint16 = 32767
                _ = v44
            } else {
                _ = len(v1)
                if true {
                    v3 = 65535
                }
            }
            _ = fi40
        }
        var v45 int32 = (v17 % (((v17 << ((65535) & 63))) | 1))
        _ = fi39
        _ = v45
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (("foo bar" + ("foo bar" + "")) > "\n\t") {
        for fi46 := 0; fi46 < 7; fi46++ {
            for fi47 := 0; fi47 < 15; fi47++ {
                v48 := GenS2{F0: "foo bar", F1: 1, F2: (false || true)}
                if false {
                    if false {
                        if (v38 == v38) {
                            var v49 float32 = 0.0
                            _ = v49
                        } else {
                            v38 += (v38 & (v38 % ((v38) | 1)))
                            v50 := GenS5{F0: false}
                            if true {
                                v51 := GenS6{F0: 9223372036854775807, F1: 100}
                                _ = v51
                            } else {
                                if true {
                                    fmt.Println(v4, v17)
                                    var v52 int = fi47
                                    _ = v52
                                }
                                v53 := GenS7{F0: 1, F1: (("\n\t" + "") + ("hello" + ("\n\t" + "")))}
                                _ = v53
                            }
                            _ = v50
                        }
                    }
                    if true {
                        var v54 int32 = (v17 - 42)
                        if (true && false) {
                            var v55 uint = (v4 >> ((65535) & 63))
                            _ = v55
                        }
                        _ = v54
                    } else {
                        v56 := GenS8{F0: 4294967295}
                        _ = v56
                    }
                    v57 := GenS5{F0: true}
                    _ = len(v5)
                    v4 = genFn0()
                    _ = v57
                }
                v58 := GenS9{F0: (v38 << ((genFn0()) & 63))}
                _ = fi47
                _ = v48
                _ = v58
            }
            fmt.Println(v38, v4, v17)
            _ = len(v2)
            _ = fi46
        }
    } else {
        for fi59 := 0; fi59 < 13; fi59++ {
            switch (fi59 << ((v4) & 63)) {
            case -92:
                var v60 bool = (!(459 != v38))
                _ = v60
            case 566:
                var v61 float64 = float64(3.14159)
                _ = v61
            default:
                for fi62 := 0; fi62 < 5; fi62++ {
                    var v63 uint32 = 2
                    _ = fi62
                    _ = v63
                }
            }
            var v64 uint32 = 4294967295
            _ = len(v6)
            _ = fi59
            _ = v64
        }
        v4 += 65535
    }
    var v65 uint = genFn0()
    fmt.Println(v3, v17)
    if true {
        switch 2147483647 {
        case 0:
            v4 -= v4
        case -964:
            var v66 uint64 = 2
            _ = v66
        case 1:
            var v67 uint64 = 4294967295
            _ = v67
        }
    }
    switch 128 {
    case -240:
        v4 = (v3 + (v3 / (((^(v65))) | 1)))
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v65
}
