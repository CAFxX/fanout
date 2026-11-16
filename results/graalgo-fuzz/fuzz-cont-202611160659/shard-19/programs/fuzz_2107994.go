package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = 42
    var v2 int32 = 1
    var v3 bool = false
    v4 := make(map[string]string)
    for fi5 := 0; fi5 < 10; fi5++ {
        v3 = v3
        var v6 int8 = 72
        v7 := GenS1{F0: 65535}
        _ = fi5
        _ = v6
        _ = v7
    }
    v3 = v3
    if (42 <= 42) {
        v4["foo bar"] = ""
        var v8 uint64 = 2
        _ = v8
    } else {
        switch "" {
        case "":
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := GenS1{F0: 835}
                for fi11 := 0; fi11 < 20; fi11++ {
                    if (-2147483648 >= (^(fi11))) {
                        var v12 float32 = float32(0.1)
                        _ = v12
                    } else {
                        v2 += 780
                    }
                    var v13 uint16 = 32767
                    _ = fi11
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
        case "foo bar":
            for fi14 := 0; fi14 < 14; fi14++ {
                if v3 {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        v16 := v4["\n\t"]
                        var v17 float32 = 3.14159
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    switch (fi14 - fi14) {
                    case 42:
                        fmt.Println(fi14)
                    case -411:
                        var v18 int32 = v2
                        _ = v18
                    case 32768:
                        fmt.Println(v2, v2, v3)
                    default:
                        switch fi14 {
                        case 387:
                            var v19 uint64 = 640
                            _ = v19
                        case -896:
                            var v20 float64 = (0.1 + ((1.5 / -2.25) * 0.1))
                            _ = v20
                        }
                    }
                    if true {
                        for fi21 := 0; fi21 < 15; fi21++ {
                            var v22 int64 = v1
                            if ((fi21 >= (fi21 ^ fi21)) && v3) {
                                switch 4294967295 {
                                case 128:
                                    if true {
                                        if (42 >= 42) {
                                            delete(v4, (("a" + ("" + "foo bar")) + "foo bar"))
                                            v22 -= v1
                                        }
                                        _ = len(v4)
                                        switch "" {
                                        case "hello":
                                            if true {
                                                var v23 int = fi14
                                                var v24 int32 = v2
                                                switch "a" {
                                                case "foo bar":
                                                    v25 := []int{(fi14 >> ((609) & 63))}
                                                    _ = v25
                                                default:
                                                    var v26 float64 = -2.25
                                                    var v27 uint32 = 2147483647
                                                    fmt.Println(fi21, v2)
                                                    fmt.Println(v2)
                                                    v28 := GenS1{F0: v27}
                                                    v3 = ((v3 && (32767 != 32767)) && true)
                                                    _ = v26
                                                    _ = v27
                                                    _ = v28
                                                }
                                                _ = v23
                                                _ = v24
                                            }
                                        case "\n\t":
                                            v1 += (^(v1))
                                        }
                                    }
                                }
                                v1 = 42
                            }
                            var v29 uint = 1
                            _ = fi21
                            _ = v22
                            _ = v29
                        }
                        var v30 uint16 = 642
                        _ = v30
                    }
                } else {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v32 := []int{(fi14 | fi14), fi31}
                        _ = fi31
                        _ = v32
                    }
                }
                var v33 int = 128
                _ = fi14
                _ = v33
            }
        }
        v3 = (!(false && (127 < -128)))
        var v34 float64 = (((0.1 * (float64((((0.1 + (((float32((3.14159 / (1e10 * (0.1 * 1.5)))) - float32((1e10 - (-2.25 * float64(((float32(1.5) + 1.5) + 1e10)))))) + float32(3.14159)) / -2.25)) - (-2.25 * (1e10 * 1e10))) + 3.14159)) * -2.25)) * float64((float32(1e10) * float32((-2.25 - (0.1 * (-2.25 * 1e10))))))) / -2.25)
        _ = v34
    }
    if (42 > 42) {
        _ = len(v4)
        v35 := v4["\n\t"]
        _ = v35
    } else {
        if v3 {
            if ((2 < 0) && (("\n\t" + "\n\t") >= "")) {
                v3 = (!(!v3))
            }
            var v36 int8 = -128
            var v37 uint8 = 127
            _ = v36
            _ = v37
        }
        var v38 uint16 = 65535
        _ = v38
    }
    v3 = true
    v2 -= (v2 ^ v2)
    fmt.Println(v3, v3)
    v39 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
