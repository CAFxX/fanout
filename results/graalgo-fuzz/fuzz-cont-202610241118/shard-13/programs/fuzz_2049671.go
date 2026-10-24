package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 bool }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 string; F1 uint; F2 int16 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint64) uint {
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(p0)
        p0 = (^(p0))
        _ = fi6
    }
    var v7 uint16 = 2
    fmt.Println(v7, p0)
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint32 = uint32(v7)
        var v10 int = (fi8 >> ((42) & 63))
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 uint16) int {
    var v11 int64 = -1
    var v12 bool = (!true)
    _ = p0
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    var v5 string = ("foo bar" + "")
    v5 = "\n\t"
    switch genFn1(109) {
    case -2147483648:
        var v13 uint32 = 65535
        _ = v13
    }
    var v14 int32 = 2147483647
    if false {
        var v15 uint = 560
        if (-1 > (v14 + (v14 >> (((v15 % (((^(v15))) | 1))) & 63)))) {
            v15 = genFn0(4294967295)
        }
        var v16 uint16 = 0
        _ = v15
        _ = v16
    } else {
        if true {
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 float32 = 1e10
                _ = fi17
                _ = v18
            }
        } else {
            for fi19 := 0; fi19 < 7; fi19++ {
                if true {
                    v14 = (v14 | 42)
                }
                if true {
                    if true {
                        fmt.Println(fi19, v5, fi19)
                        v20 := GenS1{F0: "foo bar", F1: 4294967295, F2: true}
                        _ = v20
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            v5 = "foo bar"
                        }
                        var v21 string = v5
                        switch "\n\t" {
                        case "hello":
                            var v22 int8 = 0
                            _ = v22
                        }
                        _ = v21
                    }
                    var v23 uint64 = 672
                    if (42 == v14) {
                        v5 = v5
                        for fi24 := 0; fi24 < 2; fi24++ {
                            v5 = (v5 + (v5 + ((v5 + v5) + "foo bar")))
                            _ = fi24
                        }
                        fmt.Println(v14, fi19)
                    } else {
                        var v25 uint = 874
                        v5 = "a"
                        _ = v25
                    }
                    _ = v23
                }
                switch v5 {
                case "":
                    v5 = "\n\t"
                case "\n\t":
                    var v26 float64 = (-2.25 + ((float64((1e10 / (1.5 / float32(-2.25)))) - 0.1) * (0.0 / 0.1)))
                    _ = v26
                case "a":
                    v5 = "a"
                default:
                    for fi27 := 0; fi27 < 1; fi27++ {
                        switch (-(fi27)) {
                        case 4294967295:
                            var v28 float32 = (0.0 * float32(1.5))
                            _ = v28
                        case 127:
                            var v29 float64 = 0.1
                            _ = v29
                        case 9223372036854775807:
                            var v30 float32 = 1.5
                            _ = v30
                        }
                        v14 -= v14
                        _ = fi27
                    }
                }
                _ = fi19
            }
            var v31 uint64 = 18446744073709551615
            v5 = (v5 + v5)
            _ = v31
        }
        var v32 uint = 0
        for fi33 := 0; fi33 < 15; fi33++ {
            if true {
                v34 := GenS1{F0: v5, F1: 1, F2: (!false)}
                _ = v34
            } else {
                var v35 int8 = 99
                _ = v35
            }
            v5 = v5
            _ = fi33
        }
        _ = v32
    }
    if true {
        if false {
            v36 := GenS1{F0: v5, F1: 747, F2: (!false)}
            v5 = "hello"
            _ = v36
        } else {
            var v37 int8 = 42
            if false {
                if (!true) {
                    if true {
                        delete(v1, "foo bar")
                    }
                    var v38 int64 = -1
                    _ = v38
                } else {
                    switch ("\n\t" + (v5 + "a")) {
                    case "hello":
                        v39 := GenS2{F0: v5, F1: 127, F2: v37}
                        _ = v39
                    case "a":
                        fmt.Println(v5)
                    case "":
                        v40 := GenS2{F0: v5, F1: v37, F2: (^(v37))}
                        _ = v40
                    default:
                        v14 += v14
                    }
                }
                v41 := GenS3{F0: ((219 <= 2) || (803 < 1)), F1: false}
                _ = v41
            }
            _ = v37
        }
        var v42 uint8 = 42
        v42 = (v42 >> ((18446744073709551615) & 63))
        _ = v42
    } else {
        var v43 string = v5
        switch 0 {
        case 4294967295:
            v44 := []int{42}
            _ = v44
        }
        var v45 int32 = 0
        _ = v43
        _ = v45
    }
    switch 1 {
    case -37:
        for fi46 := 0; fi46 < 13; fi46++ {
            for fi47 := 0; fi47 < 17; fi47++ {
                v48 := GenS4{F0: "a", F1: 255, F2: 610}
                _ = fi47
                _ = v48
            }
            v14 = (v14 << ((881) & 63))
            if true {
                var v49 float64 = 0.0
                for fi50 := 0; fi50 < 5; fi50++ {
                    var v51 bool = false
                    v52 := []int{fi46, (^(fi46))}
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                var v53 uint64 = 0
                _ = v49
                _ = v53
            } else {
                var v54 int64 = -689
                v14 = v14
                _ = v54
            }
            _ = fi46
        }
    case 42:
        var v55 int8 = -128
        _ = v55
    case 32767:
        v14 = (v14 * (v14 * v14))
    default:
        var v56 int8 = 42
        _ = v56
    }
    switch v5 {
    case "hello":
        var v57 float64 = (3.14159 + (0.0 + 1e10))
        _ = v57
    case "\n\t":
        v5 = (("a" + "\n\t") + v5)
    case "a":
        for fi58 := 0; fi58 < 16; fi58++ {
            var v59 uint8 = 127
            _ = fi58
            _ = v59
        }
    }
    var v60 int16 = 42
    if (((42 < v14) || (42 > 42)) && (((v5 + v5) + ("" + v5)) > v5)) {
        var v61 uint16 = 2
        for fi62 := 0; fi62 < 9; fi62++ {
            if (!(507 <= (v61 * 706))) {
                v60 = (^(v60))
            } else {
                if ((((v5 + "\n\t") + (v5 + v5)) + "\n\t") == "") {
                    v14 = (v14 - -1)
                } else {
                    v60 += v60
                }
                fmt.Println(v5, v60)
                fmt.Println(v14, v5)
            }
            _ = fi62
        }
        if true {
            var v63 int64 = 9223372036854775807
            _ = v63
        } else {
            var v64 int = 32767
            fmt.Println(v61, v61)
            v65 := GenS5{F0: v14}
            _ = v64
            _ = v65
        }
        _ = v61
    } else {
        if (0.1 == (1.5 - 1e10)) {
            v5 = (v5 + (("" + (("a" + ("a" + v5)) + v5)) + ((("\n\t" + "hello") + ((v5 + v5) + v5)) + (v5 + "a"))))
        }
    }
    v60 = 42
    v66 := v1["hello"]
    v67 := v3[("hello" + (v5 + v66))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v66
    _ = v67
}
