package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint8, p1 int16) uint64 {
    v5 := GenS1{F0: (("a" + ("a" + ("foo bar" + ""))) + ("a" + ("hello" + ("a" + ("" + "foo bar"))))), F1: 2}
    var v6 float32 = 0.1
    v7 := GenS1{F0: ((("a" + "\n\t") + ("a" + ("\n\t" + ""))) + "foo bar"), F1: 0}
    var v8 int8 = 0
    v6 -= 3.14159
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 217
}
func genFn1(p0 int, p1 string) int64 {
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(p1, p1, fi9)
        _ = fi9
    }
    p1 = ("\n\t" + "a")
    p0 += 128
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 1
    var v3 int8 = 0
    var v4 uint64 = 1
    v2 -= 0
    fmt.Println(v2, v4)
    var v10 uint = 65535
    v11 := GenS1{F0: "a", F1: 659}
    var v12 int8 = v3
    v12 = v12
    fmt.Println(v10, v2)
    var v13 uint16 = 123
    v14 := GenS1{F0: "a", F1: v13}
    v15 := GenS1{F0: ((("a" + ("foo bar" + "")) + "a") + (("a" + "") + "\n\t")), F1: uint16(v10)}
    if (!false) {
        v16 := GenS1{F0: ("\n\t" + ""), F1: v13}
        var v17 uint64 = 4294967295
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 int = (fi18 + (fi18 & fi18))
        v20 := []int{(fi18 / ((fi18) | 1)), (^(v19))}
        v21 := GenS1{F0: "a", F1: v13}
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 uint64 = (v4 % ((genFn0(v2, 1)) | 1))
    v23 := v1["hello"]
    switch "foo bar" {
    case "":
        switch -662 {
        case -801:
            v24 := GenS1{F0: "hello", F1: (^(v13))}
            _ = v24
        case 127:
            switch "hello" {
            case "\n\t":
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := GenS1{F0: "", F1: (v13 + (^(v13)))}
                    switch "hello" {
                    case "":
                        fmt.Println(v3, v4, v10)
                    }
                    _ = fi25
                    _ = v26
                }
            }
        default:
            for fi27 := 0; fi27 < 14; fi27++ {
                v2 = 127
                var v28 float64 = float64(-2.25)
                _ = fi27
                _ = v28
            }
        }
    case "hello":
        var v29 uint16 = (v13 * 0)
        _ = v29
    case "foo bar":
        v2 = v2
    default:
        for fi30 := 0; fi30 < 20; fi30++ {
            v22 -= genFn0(0, 1000)
            for fi31 := 0; fi31 < 10; fi31++ {
                switch "foo bar" {
                case "hello":
                    if (true || true) {
                        var v32 uint32 = 2147483648
                        if (fi30 == int(v10)) {
                            var v33 uint16 = 65535
                            _ = v33
                        }
                        if v23 {
                            v34 := GenS2{F0: 769}
                            fmt.Println(v23, v23, v13)
                            _ = v34
                        }
                        _ = v32
                    } else {
                        var v35 uint32 = 438
                        v36 := v1["hello"]
                        v23 = v23
                        _ = v35
                        _ = v36
                    }
                }
                _ = fi31
            }
            switch fi30 {
            case 0:
                for fi37 := 0; fi37 < 11; fi37++ {
                    var v38 uint8 = (v2 | (v2 << ((v10) & 63)))
                    _ = fi37
                    _ = v38
                }
            case 9223372036854775807:
                var v39 int = fi30
                _ = v39
            case 340:
                var v40 string = "foo bar"
                _ = v40
            default:
                for fi41 := 0; fi41 < 12; fi41++ {
                    switch (fi41 ^ fi41) {
                    case 32767:
                        fmt.Println(v4, v4, v2)
                    case 128:
                        v42 := []int{(fi41 << ((v10) & 63))}
                        _ = v42
                    case -1:
                        v43 := []int{32767}
                        _ = v43
                    }
                    _ = fi41
                }
            }
            _ = fi30
        }
    }
    if (v22 < v4) {
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
