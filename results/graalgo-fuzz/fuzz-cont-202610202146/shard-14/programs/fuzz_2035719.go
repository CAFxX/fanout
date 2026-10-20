package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 bool; F1 uint64; F2 uint8 }
func genFn0() int64 {
    var v5 int32 = 1
    var v6 uint = 255
    v5 = int32(v6)
    _ = v5
    _ = v6
    return -629
}
func genFn1(p0 uint8) int32 {
    for fi7 := 0; fi7 < 20; fi7++ {
        if false {
            fmt.Println(fi7)
        }
        v8 := GenS1{F0: fi7, F1: 255, F2: -9223372036854775808}
        v9 := GenS2{F0: 18446744073709551615, F1: 9223372036854775807}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{32768, -2147483648, 127, -454}
    if len(v10) > 0 {
        v10[0] = 32768
    }
    _ = p0
    _ = v10
    return -1
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]string)
    v3 := []string{("foo bar" + ("hello" + "a")), ""}
    v4 := make(map[string]bool)
    v1 += 2
    for fi11 := 0; fi11 < 5; fi11++ {
        if ((130 > (v1 % ((v1) | 1))) || true) {
            for fi12 := 0; fi12 < 20; fi12++ {
                v1 += v1
                var v13 uint64 = 338
                var v14 uint64 = v1
                if (false || true) {
                    if (!false) {
                        if (("foo bar" + "hello") == "\n\t") {
                            var v15 uint8 = 128
                            v16 := GenS1{F0: (fi11 - 128), F1: 0, F2: 42}
                            _ = v15
                            _ = v16
                        }
                    }
                    fmt.Println(v1, v14)
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            v17 := GenS3{F0: 2, F1: 42}
            _ = v17
        } else {
            if len(v3) > 0 {
                v3[0] = (("hello" + ((("\n\t" + "") + ((((("\n\t" + "a") + ("foo bar" + "hello")) + "\n\t") + "foo bar") + "")) + "\n\t")) + "hello")
            }
            fmt.Println(fi11)
        }
        v18 := GenS2{F0: 0, F1: 0}
        _ = fi11
        _ = v18
    }
    v1 -= (v1 / (((^(v1))) | 1))
    fmt.Println(v1)
    for fi19 := 0; fi19 < 16; fi19++ {
        v1 = (v1 >> ((4294967295) & 63))
        _ = fi19
    }
    if ((!false) && (0 > 0)) {
        var v20 float64 = float64(3.14159)
        var v21 string = "foo bar"
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 7; fi22++ {
            v1 = v1
            var v23 uint64 = (v1 / (((v1 * v1)) | 1))
            v24 := v4["foo bar"]
            _ = fi22
            _ = v23
            _ = v24
        }
        v1 -= v1
    }
    v25 := GenS4{F0: (-670 == 738), F1: (^(v1)), F2: 0}
    if (!true) {
        v1 = v1
        v1 = (v1 & v1)
    } else {
        switch -2147483648 {
        case -1:
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 uint16 = 618
                _ = fi26
                _ = v27
            }
        case 127:
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 float64 = 0.0
                var v30 float32 = (float32(v29) * float32(v29))
                switch 0 {
                case -9223372036854775808:
                    if false {
                        var v31 uint32 = 2147483648
                        v29 = (1.5 + 0.0)
                        if ((0 < 881) && ("\n\t" <= "foo bar")) {
                            var v32 float32 = (v30 * v30)
                            v31 = 65535
                            _ = v32
                        }
                        _ = v31
                    } else {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            fmt.Println(fi33, fi33, v29)
                            v34 := GenS3{F0: 207, F1: (fi28 & (fi33 * (-(fi28))))}
                            _ = fi33
                            _ = v34
                        }
                        v35 := v4["foo bar"]
                        v36 := GenS2{F0: 583, F1: -952}
                        _ = v35
                        _ = v36
                    }
                }
                _ = fi28
                _ = v29
                _ = v30
            }
        }
    }
    var v37 float64 = (0.1 * 0.0)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        var v38 int = -2147483648
        var v39 string = ((("foo bar" + "a") + ("foo bar" + "a")) + ("foo bar" + ("a" + (("\n\t" + ("hello" + ("a" + ((("a" + "foo bar") + "") + "hello")))) + ("hello" + "foo bar")))))
        _ = v38
        _ = v39
    }
    var v40 uint32 = 915
    switch 339 {
    case 4294967295:
        var v41 int16 = 0
        _ = v41
    }
    v42 := GenS3{F0: 18446744073709551615, F1: 2147483647}
    var v43 uint64 = v1
    switch ("hello" + ("foo bar" + ("\n\t" + ("\n\t" + "")))) {
    case "hello":
        v37 -= 1e10
    case "foo bar":
        for fi44 := 0; fi44 < 2; fi44++ {
            _ = len(v4)
            var v45 float32 = -2.25
            fmt.Println(v40, v43)
            _ = fi44
            _ = v45
        }
    }
    v43 = (v1 ^ v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v42
    _ = v43
}
