package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint16 }
func genFn0(p0 uint, p1 uint, p2 uint) string {
    switch 32767 {
    case 127:
        switch ("\n\t" + ("hello" + ("foo bar" + (("" + ("hello" + "\n\t")) + ("" + ("hello" + "\n\t")))))) {
        case "hello":
            p2 -= p0
        default:
            var v7 int16 = -32768
            _ = v7
        }
    }
    if false {
        switch "foo bar" {
        case "foo bar":
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 bool = (42 == 466)
                _ = fi8
                _ = v9
            }
        case "a":
            fmt.Println(p0, p1)
        default:
            for fi10 := 0; fi10 < 8; fi10++ {
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 bool = ((-32768 > 32767) && ((fi11 ^ (-(fi10))) < fi10))
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                switch fi13 {
                case 32768:
                    v15 := []int{fi14, (fi13 + fi13), (fi13 << ((18446744073709551615) & 63)), (^(fi13))}
                    _ = v15
                }
                _ = fi14
            }
            var v16 int8 = 127
            _ = fi13
            _ = v16
        }
        v17 := []int{}
        _ = v17
    } else {
        var v18 int8 = 100
        var v19 int8 = v18
        v20 := GenS1{F0: 111, F1: 1, F2: 282}
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 int8 = -108
    for fi22 := 0; fi22 < 8; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            v21 = v21
            _ = fi23
        }
        if (true && false) {
            switch (fi22 + fi22) {
            case -639:
                fmt.Println(fi22)
            case 32768:
                p2 -= (^(p1))
            }
            var v24 uint = p1
            if true {
                var v25 uint64 = 0
                _ = v25
            }
            _ = v24
        } else {
            for fi26 := 0; fi26 < 19; fi26++ {
                p0 += p2
                if (65535 <= 0) {
                    var v27 int16 = 730
                    if true {
                        var v28 uint16 = 983
                        if (false && true) {
                            var v29 bool = true
                            v30 := []int{int(v27), (fi22 & (fi26 % ((fi22) | 1)))}
                            _ = v29
                            _ = v30
                        } else {
                            var v31 uint8 = 255
                            _ = v31
                        }
                        if (879 == v28) {
                            var v32 int32 = 0
                            _ = v32
                        }
                        _ = v28
                    } else {
                        var v33 uint32 = 4294967295
                        _ = v33
                    }
                    _ = v27
                } else {
                    p0 -= 0
                    v34 := GenS1{F0: (-(fi26)), F1: -2147483648, F2: 1}
                    _ = v34
                }
                _ = fi26
            }
        }
        v35 := GenS1{F0: fi22, F1: 0, F2: 23}
        _ = fi22
        _ = v35
    }
    switch (((("\n\t" + "\n\t") + "\n\t") + "foo bar") + ("a" + "hello")) {
    case "a":
        v21 += (v21 & (v21 % ((v21) | 1)))
    default:
        v36 := GenS1{F0: 4294967295, F1: -2147483648, F2: 277}
        _ = v36
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return "hello"
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint64 = 42
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    v6 := make(map[int]uint32)
    for fi37 := 0; fi37 < 15; fi37++ {
        v38 := v6[v1]
        _ = fi37
        _ = v38
    }
    var v39 uint16 = 255
    v1 = v1
    v40 := v4[((("\n\t" + ("" + "foo bar")) + "\n\t") + "a")]
    var v41 float64 = 3.14159
    var v42 uint = 770
    v43 := GenS1{F0: v1, F1: -513, F2: uint16(v2)}
    var v44 int64 = 42
    v45 := GenS1{F0: (v1 >> ((2) & 63)), F1: 979, F2: (v39 ^ 42)}
    fmt.Println(v44, v41)
    if (false && (v40 && (22 <= 4))) {
        var v46 int8 = 42
        _ = v46
    }
    v3[v1] = "hello"
    var v47 uint16 = (v39 / (((v39 ^ 0)) | 1))
    var v48 uint64 = 840
    var v49 float32 = 1.5
    delete(v5, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
