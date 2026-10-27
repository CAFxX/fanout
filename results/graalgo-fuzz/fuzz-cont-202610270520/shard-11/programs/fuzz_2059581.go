package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint8 }
func genFn0(p0 uint32, p1 string) int16 {
    var v4 float64 = -2.25
    var v5 float64 = float64(3.14159)
    v5 -= 1e10
    v6 := GenS1{F0: -32768, F1: 0, F2: 42}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int8 = 96
    v2 := []string{"", "a", (("\n\t" + "foo bar") + "\n\t"), "a"}
    var v3 uint16 = 255
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(fi7, v1)
        v2 = append(v2, ("hello" + "hello"))
        _ = fi7
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        if (2 <= 236) {
            v1 += (v1 - (v1 ^ -24))
        }
        switch (fi8 & fi8) {
        case -1:
            var v9 int16 = -722
            _ = v9
        case -2147483648:
            if false {
                v3 = (v3 + v3)
            }
        }
        _ = fi8
    }
    v1 = (v1 + (v1 | v1))
    var v10 int8 = (v1 * (v1 >> ((2) & 63)))
    if len(v2) > 0 {
        v2[0] = ("" + "hello")
    }
    if len(v2) > 0 {
        v2[0] = ("foo bar" + "a")
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        if false {
            switch fi11 {
            case 128:
                for fi12 := 0; fi12 < 14; fi12++ {
                    switch fi11 {
                    case -778:
                        v1 = 42
                    }
                    _ = len(v2)
                    v10 = v10
                    var v13 bool = true
                    _ = fi12
                    _ = v13
                }
            }
            var v14 int32 = 42
            v15 := GenS1{F0: 0, F1: v1, F2: 147}
            _ = v14
            _ = v15
        }
        var v16 float32 = float32((float64(1.5) / 1.0))
        if (!(false && false)) {
            if false {
                v2 = append(v2, "")
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            }
            v3 -= v3
        }
        _ = fi11
        _ = v16
    }
    switch ("hello" + ("foo bar" + "foo bar")) {
    case "":
        if false {
            var v17 int8 = (v1 >> ((77) & 63))
            if (!true) {
                v17 += 0
                v10 = (v17 << ((746) & 63))
                for fi18 := 0; fi18 < 17; fi18++ {
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    v1 = (-(v17))
                    switch int(v3) {
                    case -603:
                        var v19 float32 = 1.5
                        _ = v19
                    case 0:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -1:
                        v3 = v3
                    }
                    _ = fi18
                }
            } else {
                if false {
                    v2 = append(v2, "foo bar")
                }
            }
            _ = v17
        } else {
            for fi20 := 0; fi20 < 14; fi20++ {
                switch fi20 {
                case -597:
                    var v21 uint32 = 2147483647
                    _ = v21
                case 42:
                    v3 += v3
                case 1:
                    v3 -= (v3 >> ((255) & 63))
                }
                _ = fi20
            }
            var v22 float64 = -2.25
            _ = v22
        }
    case "\n\t":
        if (!(v3 != (v3 + v3))) {
            v10 = (v1 & (v10 + (-(v10))))
            if ((!true) || (0.0 != 3.14159)) {
                v23 := GenS1{F0: 42, F1: (v10 << ((980) & 63)), F2: 0}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v23
            } else {
                var v24 float64 = float64(1e10)
                v10 -= v10
                _ = v24
            }
            var v25 uint8 = 2
            _ = v25
        }
    default:
        if (805 <= 2) {
            var v26 uint8 = 127
            _ = v26
        }
    }
    switch ((("\n\t" + "a") + "") + "hello") {
    case "":
        _ = len(v2)
    }
    var v27 int = -2147483648
    v2 = append(v2, "a")
    fmt.Println(v10, v3, v27)
    var v28 int8 = (-(v10))
    switch 128 {
    case 32767:
        for fi29 := 0; fi29 < 20; fi29++ {
            var v30 int = v27
            _ = fi29
            _ = v30
        }
    case 32768:
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 int8 = v10
            fmt.Println(v1, fi31, v27)
            _ = fi31
            _ = v32
        }
    case -9223372036854775808:
        v1 = (v28 % ((v28) | 1))
    }
    var v33 int8 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
}
