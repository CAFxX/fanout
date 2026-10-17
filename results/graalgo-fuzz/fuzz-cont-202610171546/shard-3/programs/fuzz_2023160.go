package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint64; F2 string }
func genFn0(p0 int16, p1 uint16, p2 string) int8 {
    v7 := []int{-133, 9223372036854775807, 925, -1}
    if (!("foo bar" < ((p2 + p2) + "a"))) {
        var v8 int = 127
        _ = v8
    } else {
        var v9 bool = true
        v9 = (p2 <= p2)
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        if false {
            var v11 uint8 = 2
            switch ("hello" + (("hello" + ("\n\t" + p2)) + (((("foo bar" + p2) + "foo bar") + (p2 + (p2 + ((((p2 + p2) + "\n\t") + "a") + p2)))) + p2))) {
            case "\n\t":
                for fi12 := 0; fi12 < 13; fi12++ {
                    if true {
                        var v13 int8 = -13
                        _ = v13
                    } else {
                        fmt.Println(p2, p0, v11)
                        fmt.Println(fi12, fi12)
                        v14 := GenS1{F0: 0}
                        _ = v14
                    }
                    _ = fi12
                }
            case "foo bar":
                p2 = p2
            }
            _ = v11
        }
        var v15 int64 = -846
        fmt.Println(p1, p2)
        _ = fi10
        _ = v15
    }
    p0 = p0
    if false {
        p1 -= 1
        if (false && (3.14159 <= (((1.5 / 1.0) - (float32(-2.25) + float32(float64(1.5)))) / 0.1))) {
            if (-9223372036854775808 < 542) {
                p0 = (p0 * p0)
                if (1 >= 127) {
                    if false {
                        for fi16 := 0; fi16 < 17; fi16++ {
                            var v17 string = p2
                            v7 = append(v7, fi16)
                            _ = fi16
                            _ = v17
                        }
                        v18 := GenS2{F0: 0}
                        fmt.Println(p1, p1)
                        _ = v18
                    }
                    fmt.Println(p1, p0, p2)
                }
            }
            if len(v7) > 0 {
                v7[0] = 32767
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 int = (fi19 + (fi19 << ((1) & 63)))
                for fi21 := 0; fi21 < 3; fi21++ {
                    if len(v7) > 0 {
                        v7[0] = fi19
                    }
                    p1 = (p1 / ((p1) | 1))
                    _ = fi21
                }
                var v22 uint16 = (^(p1))
                _ = fi19
                _ = v20
                _ = v22
            }
        } else {
            p2 = (p2 + "hello")
        }
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 127
}
func genFn1(p0 bool) uint8 {
    var v23 uint16 = 1
    if (0 <= 0) {
        v23 -= v23
        v24 := GenS1{F0: 4294967295}
        v25 := []int{-2147483648, 622, 4294967295}
        _ = v24
        _ = v25
    } else {
        if (1e10 <= float32(1.5)) {
            var v26 float32 = 1.5
            var v27 string = ""
            if (!p0) {
                if (true || (((0 > -2147483648) || p0) && (-2147483648 <= -9223372036854775808))) {
                    fmt.Println(v26)
                    v28 := GenS1{F0: 2147483648}
                    switch ("a" + (v27 + (("a" + v27) + (v27 + ((v27 + ("\n\t" + ((v27 + v27) + "\n\t"))) + "a"))))) {
                    case "\n\t":
                        for fi29 := 0; fi29 < 12; fi29++ {
                            v26 += 0.0
                            var v30 bool = (-1 != -128)
                            _ = fi29
                            _ = v30
                        }
                    case "foo bar":
                        v31 := []int{}
                        _ = v31
                    case "a":
                        for fi32 := 0; fi32 < 2; fi32++ {
                            v33 := GenS2{F0: 255}
                            var v34 int16 = -1
                            v35 := GenS3{F0: v23, F1: 21, F2: 0}
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    }
                    _ = v28
                }
            }
            _ = v26
            _ = v27
        }
        v23 -= 42
        fmt.Println(v23, p0)
    }
    _ = p0
    _ = v23
    return 128
}

func main() {
    var v1 uint = 1
    var v2 int16 = -1
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v5 bool = (799 <= 255)
    v6 := make(map[uint32]uint32)
    if (!false) {
        v2 = v2
        v36 := GenS4{F0: v5, F1: 2, F2: ("hello" + "a")}
        v37 := GenS1{F0: 510}
        _ = v36
        _ = v37
    } else {
        v1 -= v1
    }
    var v38 int8 = 127
    if ((!true) || (835 != 1)) {
        var v39 float32 = (float32(1.5) + float32(3.14159))
        var v40 uint32 = uint32(v1)
        _ = v39
        _ = v40
    } else {
        for fi41 := 0; fi41 < 1; fi41++ {
            v2 += v2
            if v5 {
                v3["foo bar"] = 928
            }
            _ = fi41
        }
        v2 += (v2 + -884)
        delete(v3, "foo bar")
    }
    var v42 int16 = 1000
    var v43 float32 = 1.5
    if (!("foo bar" > (("hello" + "hello") + "hello"))) {
        if (0 >= 1) {
            switch 42 {
            case 32768:
                fmt.Println(v2, v38)
            case 2147483647:
                var v44 int64 = 9223372036854775807
                _ = v44
            case -6:
                v45 := GenS1{F0: 2147483648}
                _ = v45
            }
            var v46 uint8 = 1
            v38 = (v38 << ((uint(v46)) & 63))
            _ = v46
        } else {
            var v47 int8 = v38
            _ = v47
        }
        delete(v4, -1)
    }
    if v5 {
        if (genFn1(true) > 128) {
            var v48 uint64 = 1
            switch ("foo bar" + ("\n\t" + (("a" + "foo bar") + "\n\t"))) {
            case "hello":
                v49 := GenS2{F0: v1}
                _ = v49
            case "foo bar":
                v5 = (((1e10 + 1e10) + (-2.25 * 1.5)) >= 1e10)
            case "":
                v38 -= v38
            }
            var v50 uint8 = 255
            _ = v48
            _ = v50
        } else {
            fmt.Println(v1, v2)
            v2 += -32768
        }
        var v51 int64 = 42
        _ = v51
    } else {
        var v52 int8 = 0
        var v53 int16 = 42
        _ = v52
        _ = v53
    }
    v2 += v2
    v38 -= genFn0(v2, 42, "hello")
    v43 += 0.0
    for fi54 := 0; fi54 < 10; fi54++ {
        switch "foo bar" {
        case "a":
            var v55 uint32 = 0
            _ = v55
        }
        _ = fi54
    }
    v43 += 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
