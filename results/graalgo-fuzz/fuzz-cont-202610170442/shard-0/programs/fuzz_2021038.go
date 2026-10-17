package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint32) uint64 {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float32 = 0.1
        var v9 float32 = ((v8 * -2.25) / float32(-2.25))
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = p0
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: 0}
        p0 = p0
        if false {
            fmt.Println(fi10, p0)
        } else {
            switch "a" {
            case "\n\t":
                v12 := GenS1{F0: -1}
                _ = v12
            case "":
                p0 -= p0
            }
            v13 := GenS1{F0: 0}
            if false {
                fmt.Println(fi10, p0)
            }
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v14 := GenS2{F0: 756}
    var v15 bool = false
    _ = p0
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]bool)
    fmt.Println(v1)
    if false {
        var v16 uint = 536
        switch (("a" + "foo bar") + "a") {
        case "hello":
            if true {
                var v17 string = (("foo bar" + ("" + "hello")) + "foo bar")
                switch -923 {
                case -310:
                    v18 := GenS2{F0: 184}
                    _ = v18
                case 4294967295:
                    fmt.Println(v16, v1)
                case -822:
                    var v19 uint16 = 260
                    _ = v19
                }
                var v20 int16 = -1
                _ = v17
                _ = v20
            }
        case "\n\t":
            for fi21 := 0; fi21 < 11; fi21++ {
                if ((!false) || (2 != (^(v16)))) {
                    v22 := []int{}
                    var v23 int16 = 32767
                    for fi24 := 0; fi24 < 17; fi24++ {
                        var v25 int = fi21
                        switch "foo bar" {
                        case "hello":
                            var v26 int32 = 42
                            _ = v26
                        default:
                            var v27 int = 32767
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                    }
                    _ = v22
                    _ = v23
                }
                for fi28 := 0; fi28 < 3; fi28++ {
                    var v29 string = ""
                    _ = fi28
                    _ = v29
                }
                _ = fi21
            }
        }
        if ((true && ((v1 > v1) || (false || false))) || ((!(((-1 <= 32767) && false) && (false && (-1 != -49)))) && true)) {
            v30 := GenS3{F0: false, F1: ""}
            v31 := []int{-567, 9223372036854775807}
            _ = v30
            _ = v31
        }
        _ = v16
    } else {
        v1 += v1
        switch 126 {
        case 127:
            var v32 string = ""
            _ = v32
        case 2147483647:
            v33 := []int{-691}
            _ = v33
        default:
            v34 := GenS4{F0: 852}
            _ = v34
        }
        v35 := GenS1{F0: 32767}
        _ = v35
    }
    if (!(!false)) {
        v3[32767] = "a"
        v36 := []int{32768, -9223372036854775808, 292}
        _ = v36
    } else {
        switch "foo bar" {
        case "foo bar":
            var v37 int = 42
            _ = v37
        }
        v1 += 0
        v38 := GenS4{F0: 42}
        _ = v38
    }
    v1 -= -1
    var v39 float32 = -2.25
    v39 += 3.14159
    for fi40 := 0; fi40 < 5; fi40++ {
        v41 := []int{-884, (fi40 % ((fi40) | 1))}
        v1 = 9223372036854775807
        var v42 uint16 = 2
        _ = fi40
        _ = v41
        _ = v42
    }
    switch "foo bar" {
    case "\n\t":
        v1 -= (v1 % ((v1) | 1))
    }
    for fi43 := 0; fi43 < 10; fi43++ {
        var v44 uint16 = 32767
        v1 = (v1 & (v1 | (-(v1))))
        _ = fi43
        _ = v44
    }
    v45 := GenS1{F0: 42}
    switch "hello" {
    case "foo bar":
        v39 = -2.25
    default:
        if false {
            var v46 uint32 = 681
            _ = v46
        } else {
            var v47 uint16 = 65535
            _ = v47
        }
    }
    v48 := GenS5{F0: false, F1: 2147483647, F2: 2}
    if true {
        v49 := []int{74, 128, 32767, 4294967295}
        for fi50 := 0; fi50 < 19; fi50++ {
            var v51 uint64 = 333
            var v52 float64 = (((-2.25 / 1.5) - 0.1) / (float64((0.1 / (0.1 / (v39 - 1e10)))) / float64((v39 - v39))))
            fmt.Println(v1, v51)
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = v49
    }
    if false {
        var v53 uint = 255
        var v54 int8 = 64
        v55 := GenS6{F0: 2}
        _ = v53
        _ = v54
        _ = v55
    } else {
        var v56 string = ("foo bar" + "")
        v56 = "\n\t"
        if (807 == 1) {
            v57 := GenS4{F0: -1}
            switch ("foo bar" + v56) {
            case "":
                if true {
                    var v58 uint16 = 404
                    if (v56 != v56) {
                        v56 = v56
                        switch (v56 + (v56 + "foo bar")) {
                        case "\n\t":
                            if (!true) {
                                var v59 int = -1
                                v58 = 0
                                _ = v59
                            }
                        }
                        fmt.Println(v39, v39)
                    } else {
                        v60 := v2[1]
                        _ = v60
                    }
                    _ = v58
                }
            case "\n\t":
                v1 += v1
            case "foo bar":
                v1 = (v1 >> ((4294967295) & 63))
            }
            switch v56 {
            case "foo bar":
                var v61 float64 = 0.0
                _ = v61
            }
            _ = v57
        } else {
            v62 := []int{127, int(v1), 1}
            fmt.Println(v56, v39)
            _ = v62
        }
        _ = v56
    }
    fmt.Println(v39)
    var v63 string = ("a" + (("hello" + "hello") + "foo bar"))
    var v64 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v6
    _ = v63
    _ = v64
}
