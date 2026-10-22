package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint16; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 int32; F2 int32 }
func genFn0(p0 int32, p1 string) int16 {
    p1 = ("\n\t" + p1)
    switch p1 {
    case "":
        var v7 int64 = 42
        _ = v7
    case "foo bar":
        v8 := []int{-2147483648, 0}
        _ = v8
    }
    p1 = (p1 + (p1 + "\n\t"))
    v9 := GenS1{F0: -1, F1: 190, F2: 255}
    _ = p0
    _ = p1
    _ = v9
    return 1000
}
func genFn1(p0 int, p1 string) uint16 {
    p1 = p1
    var v10 uint64 = 1
    v11 := []int{(^(p0))}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 1
    v4 := []string{((("a" + "\n\t") + "\n\t") + "\n\t"), "\n\t", "a"}
    v5 := make(map[string]bool)
    var v6 bool = (!true)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 += v3
    v6 = (!((0 == 1000) || (v6 && (-2147483648 == -9223372036854775808))))
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(fi12, v6)
        if (fi12 >= fi12) {
            var v13 uint = 65535
            fmt.Println(v13, v3)
            _ = v13
        }
        var v14 int8 = 1
        _ = fi12
        _ = v14
    }
    switch -1 {
    case 0:
        v6 = true
    case 32768:
        _ = len(v1)
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint32 = 232
        if (v6 && (v6 && (!((!v6) || true)))) {
            for fi17 := 0; fi17 < 14; fi17++ {
                if ((^(v16)) < (v16 << ((65535) & 63))) {
                    fmt.Println(fi15)
                } else {
                    v16 = (v16 % ((v16) | 1))
                }
                _ = len(v4)
                _ = fi17
            }
            v18 := GenS2{F0: 2}
            switch -2147483648 {
            case -516:
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch (fi19 >> ((1) & 63)) {
                    case -2147483648:
                        v16 -= 1
                    case 791:
                        switch "a" {
                        case "foo bar":
                            var v20 uint32 = v16
                            _ = v20
                        case "hello":
                            v21 := GenS3{F0: (^(fi19)), F1: 389, F2: 42}
                            _ = v21
                        }
                    case 4294967295:
                        switch fi19 {
                        case -1:
                            var v22 bool = (0 >= 161)
                            _ = v22
                        }
                    default:
                        v23 := GenS3{F0: fi19, F1: v3, F2: -1}
                        _ = v23
                    }
                    v24 := GenS4{F0: 0}
                    _ = fi19
                    _ = v24
                }
            case 851:
                for fi25 := 0; fi25 < 9; fi25++ {
                    if (("foo bar" + (("a" + "\n\t") + ("foo bar" + ("" + ("hello" + "hello"))))) > ("" + "foo bar")) {
                        var v26 uint16 = (v3 & 1)
                        fmt.Println(fi25)
                        v3 = (v3 ^ genFn1(fi15, ("\n\t" + ("\n\t" + (("foo bar" + "foo bar") + "foo bar")))))
                        _ = v26
                    }
                    _ = len(v4)
                    _ = fi25
                }
            case 571:
                var v27 uint8 = 42
                _ = v27
            default:
                var v28 uint16 = (^(v3))
                _ = v28
            }
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v3 = (v3 >> ((18446744073709551615) & 63))
    v3 = genFn1(127, "\n\t")
    v6 = false
    if ((v3 + (^(v3))) >= (v3 * v3)) {
        var v29 int = -2147483648
        _ = v29
    }
    v30 := GenS5{F0: 255}
    v31 := GenS3{F0: -2147483648, F1: (v3 | (v3 + 255)), F2: 1}
    v32 := GenS6{F0: v6, F1: -1, F2: 1}
    for fi33 := 0; fi33 < 18; fi33++ {
        switch fi33 {
        case -707:
            v3 = (v3 >> ((347) & 63))
        default:
            switch ("foo bar" + ("\n\t" + "a")) {
            case "a":
                var v34 bool = (!v6)
                _ = v34
            case "":
                fmt.Println(v3)
            }
        }
        v3 = (v3 & (v3 % ((v3) | 1)))
        for fi35 := 0; fi35 < 5; fi35++ {
            fmt.Println(v6, v6)
            if false {
                if (!(!(997 < 65535))) {
                    v36 := GenS6{F0: v6, F1: 0, F2: 28}
                    if (v6 || (true && (!((0.0 * 0.1) != -2.25)))) {
                        v6 = v6
                        var v37 float64 = 0.1
                        switch fi35 {
                        case 1:
                            var v38 uint64 = 1
                            _ = v38
                        }
                        _ = v37
                    }
                    _ = v36
                }
                fmt.Println(v6, v3, v6)
            }
            _ = fi35
        }
        _ = fi33
    }
    if false {
        var v39 bool = true
        v6 = ((0 != 1) && (!v6))
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
