package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int8; F1 string; F2 uint64 }
type GenS7 struct { F0 uint; F1 int; F2 bool }
func genFn0() uint64 {
    var v5 int32 = -499
    if (2 < 42) {
        switch -1 {
        case 128:
            if true {
                fmt.Println(v5)
                for fi6 := 0; fi6 < 13; fi6++ {
                    for fi7 := 0; fi7 < 19; fi7++ {
                        fmt.Println(v5, fi7)
                        v8 := GenS1{F0: 429, F1: 863}
                        _ = fi7
                        _ = v8
                    }
                    v9 := GenS1{F0: 1, F1: 1}
                    _ = fi6
                    _ = v9
                }
                v10 := GenS1{F0: -285, F1: 32767}
                _ = v10
            }
        case 127:
            v11 := GenS1{F0: -9223372036854775808, F1: 42}
            _ = v11
        case -899:
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := []int{(fi12 % (((^(fi12))) | 1)), (fi12 >> ((65535) & 63))}
                var v14 bool = true
                _ = fi12
                _ = v13
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := []int{(fi15 % ((fi15) | 1))}
                _ = fi15
                _ = v16
            }
        }
    }
    fmt.Println(v5)
    v17 := []int{1}
    var v18 float64 = (0.0 + 0.1)
    _ = v17
    _ = v18
    _ = v5
    return 0
}
func genFn1(p0 int16) uint16 {
    switch -9223372036854775808 {
    case 127:
        var v19 int16 = p0
        _ = v19
    }
    p0 = 1
    _ = p0
    return 32767
}

func main() {
    v1 := []string{"a", "a", "foo bar", "a"}
    v2 := []int{32767, 2147483647, 32767, -9223372036854775808}
    v3 := []string{("a" + (("\n\t" + ("hello" + "foo bar")) + "hello"))}
    var v4 uint = 65535
    var v20 uint32 = 2147483648
    for fi21 := 0; fi21 < 16; fi21++ {
        if (197 > 0) {
            v22 := GenS1{F0: 388, F1: -434}
            if len(v2) > 0 {
                v2[0] = (fi21 - (fi21 >> (((v4 ^ 18446744073709551615)) & 63)))
            }
            v3 = append(v3, "a")
            _ = v22
        } else {
            v23 := GenS2{F0: v20, F1: -1, F2: 0}
            _ = v23
        }
        fmt.Println(fi21, v20, v20)
        if (false || (!true)) {
            v24 := GenS1{F0: -739, F1: 42}
            _ = v24
        }
        _ = fi21
    }
    v4 -= 18446744073709551615
    v25 := GenS2{F0: (^(v20)), F1: 1, F2: v20}
    var v26 uint8 = 0
    v27 := GenS3{F0: false}
    v28 := GenS3{F0: (("" + "a") == ("\n\t" + "foo bar"))}
    v29 := GenS3{F0: true}
    v4 = v4
    for fi30 := 0; fi30 < 16; fi30++ {
        switch (fi30 % ((-670) | 1)) {
        case 128:
            var v31 uint16 = 255
            _ = v31
        case 165:
            if true {
                v32 := GenS3{F0: (0 > 42)}
                v20 += v20
                _ = v32
            } else {
                v26 += (v26 - 0)
                for fi33 := 0; fi33 < 19; fi33++ {
                    var v34 string = (("foo bar" + "\n\t") + "a")
                    v4 -= (v4 / (((^(v4))) | 1))
                    switch v34 {
                    case "hello":
                        fmt.Println(v26, v26)
                    default:
                        v34 = (v34 + (v34 + v34))
                    }
                    _ = fi33
                    _ = v34
                }
                v2 = append(v2, int(v20))
            }
        case 32768:
            if len(v1) > 0 {
                v1[0] = ("foo bar" + (("a" + ("a" + "")) + ("foo bar" + "\n\t")))
            }
        }
        v20 = (v20 | (v20 & v20))
        var v35 uint = v4
        _ = fi30
        _ = v35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        switch "a" {
        case "a":
            fmt.Println(v26)
        case "\n\t":
            v37 := GenS4{F0: v20}
            _ = v37
        case "hello":
            if (("a" <= "a") || false) {
                switch (fi36 ^ int(v20)) {
                case -9223372036854775808:
                    if (false || true) {
                        var v38 float64 = (1.5 / 3.14159)
                        _ = v38
                    }
                default:
                    _ = len(v1)
                }
                v26 -= v26
            }
        default:
            switch "foo bar" {
            case "hello":
                _ = len(v1)
            case "foo bar":
                switch "hello" {
                case "\n\t":
                    v39 := GenS5{F0: -2147483648, F1: "foo bar"}
                    _ = v39
                case "foo bar":
                    v40 := GenS2{F0: (^(v20)), F1: 2147483647, F2: v20}
                    _ = v40
                case "a":
                    fmt.Println(v26, v20)
                default:
                    fmt.Println(fi36, v26)
                }
            case "a":
                if false {
                    for fi41 := 0; fi41 < 15; fi41++ {
                        v42 := GenS6{F0: 100, F1: "foo bar", F2: 1}
                        var v43 uint16 = 1
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v44 int = (fi36 | fi41)
                        _ = fi41
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                } else {
                    fmt.Println(fi36, fi36, v20)
                }
            default:
                var v45 int32 = -1
                _ = v45
            }
        }
        _ = fi36
    }
    fmt.Println(v26, v26)
    if false {
        for fi46 := 0; fi46 < 16; fi46++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi46
        }
    } else {
        var v47 int8 = -29
        var v48 int16 = 1000
        for fi49 := 0; fi49 < 9; fi49++ {
            v26 = (v26 | 42)
            for fi50 := 0; fi50 < 7; fi50++ {
                if len(v1) > 0 {
                    v1[0] = (("" + "a") + "a")
                }
                fmt.Println(v26, v20, v4)
                var v51 int8 = v47
                v47 = (v51 % ((42) | 1))
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                _ = fi50
                _ = v51
            }
            for fi52 := 0; fi52 < 19; fi52++ {
                fmt.Println(v47)
                _ = fi52
            }
            _ = fi49
        }
        _ = v47
        _ = v48
    }
    fmt.Println(v20, v4)
    var v53 int16 = 42
    if len(v3) > 0 {
        v3[0] = ("" + ("hello" + "\n\t"))
    }
    v53 -= v53
    v54 := GenS7{F0: 42, F1: 1, F2: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v53
    _ = v54
}
