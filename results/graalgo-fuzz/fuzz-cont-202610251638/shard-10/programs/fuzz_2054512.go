package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 string; F1 bool; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint32 }
func genFn0(p0 string, p1 string, p2 uint8) uint8 {
    fmt.Println(p1, p2, p1)
    var v5 float64 = 1.5
    v6 := GenS1{F0: -525}
    var v7 float64 = v5
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 85
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 4294967295
    var v3 float64 = (3.14159 + -2.25)
    var v4 int8 = 18
    switch 0 {
    case -2147483648:
        if (true && false) {
            var v9 uint64 = 34
            v4 -= v4
            _ = v9
        }
    case 91:
        v3 += 3.14159
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        if (0.0 != (-2.25 * 0.1)) {
            v4 = 127
            for fi11 := 0; fi11 < 19; fi11++ {
                if false {
                    v12 := GenS1{F0: 0}
                    _ = v12
                }
                _ = fi11
            }
            switch (("foo bar" + "hello") + "") {
            case "foo bar":
                v13 := v1[128]
                _ = v13
            case "a":
                var v14 uint16 = 1
                _ = v14
            case "hello":
                switch "" {
                case "":
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v3 -= -2.25
                        if true {
                            v16 := []int{1, (fi15 * fi10), -2147483648, fi10}
                            switch fi10 {
                            case 42:
                                _ = len(v16)
                            default:
                                if len(v16) > 0 {
                                    v16[0] = fi15
                                }
                            }
                            _ = v16
                        }
                        var v17 int = 32768
                        _ = fi15
                        _ = v17
                    }
                }
            }
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            switch "a" {
            case "hello":
                fmt.Println(fi18, v3, fi10)
            }
            switch fi18 {
            case 128:
                var v19 uint8 = 0
                _ = v19
            case 2147483647:
                if false {
                    if (!(!false)) {
                        v20 := GenS2{F0: "a", F1: true}
                        _ = v20
                    }
                    v21 := GenS1{F0: int32(fi10)}
                    fmt.Println(fi10, fi18)
                    v22 := v1[(fi18 + 32767)]
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS2{F0: "foo bar", F1: ("foo bar" < ("a" + "a"))}
                    if (!false) {
                        var v24 float32 = ((0.1 / 1e10) - (0.1 - -2.25))
                        fmt.Println(v3, v24)
                        _ = v24
                    } else {
                        v25 := []int{(^(fi18)), fi10, int(v4)}
                        _ = len(v25)
                        _ = v25
                    }
                    _ = v23
                }
            default:
                v26 := GenS3{F0: 1000, F1: 42}
                _ = v26
            }
            _ = fi18
        }
        if (!(255 != 2)) {
            switch "a" {
            case "a":
                for fi27 := 0; fi27 < 15; fi27++ {
                    v28 := v1[(fi27 << ((65535) & 63))]
                    if (0.1 <= v3) {
                        fmt.Println(fi10)
                        var v29 bool = (((!(42 > 194)) || (396 > 255)) || v28)
                        fmt.Println(fi10, v2)
                        var v30 uint16 = 134
                        var v31 uint = 1
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    v32 := []int{}
                    _ = fi27
                    _ = v28
                    _ = v32
                }
            case "\n\t":
                v33 := GenS2{F0: ("a" + "\n\t"), F1: false}
                _ = v33
            }
        }
        _ = fi10
    }
    var v34 int16 = 897
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 uint = 218
        fmt.Println(v34, v4)
        fmt.Println(v36, v2)
        _ = fi35
        _ = v36
    }
    if (false || false) {
        v37 := []int{32768, -2147483648, 42}
        fmt.Println(v2)
        _ = v37
    } else {
        switch "" {
        case "":
            v3 -= v3
        }
        switch -9223372036854775808 {
        case 1:
            v34 += 1000
        case 128:
            for fi38 := 0; fi38 < 8; fi38++ {
                v39 := []int{(fi38 * (fi38 >> ((4294967295) & 63)))}
                _ = fi38
                _ = v39
            }
        default:
            if false {
                if (false && false) {
                    fmt.Println(v4, v3, v34)
                    fmt.Println(v3, v2, v2)
                }
            } else {
                v40 := []int{4294967295, -9223372036854775808}
                _ = v40
            }
        }
        v41 := v1[-135]
        _ = v41
    }
    var v42 uint64 = 4294967295
    v43 := GenS3{F0: (v34 - -378), F1: 42}
    v44 := []int{592, 32768, -1, 666}
    v45 := GenS4{F0: 0, F1: int16(v2)}
    fmt.Println(v3, v3, v4)
    var v46 int = 32767
    fmt.Println(v46, v4)
    if false {
        var v47 int = v46
        _ = v47
    }
    switch "hello" {
    case "":
        v48 := GenS5{F0: ("a" + ("a" + ("\n\t" + ((("" + "a") + ("a" + "")) + "foo bar")))), F1: (v2 > v2), F2: v4}
        _ = v48
    default:
        switch ("hello" + (("" + (("\n\t" + ("hello" + "foo bar")) + ("\n\t" + "hello"))) + "")) {
        case "a":
            v49 := GenS1{F0: -802}
            _ = v49
        default:
            v4 -= v4
        }
    }
    v50 := GenS6{F0: 1, F1: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v50
}
