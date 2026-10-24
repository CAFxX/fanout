package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    v3 := []int{0, -278, 128}
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    switch 0 {
    case 42:
        var v6 float32 = 0.1
        _ = v6
    }
    if (42 <= -10) {
        if len(v3) > 0 {
            v3[0] = 0
        }
        if true {
            if false {
                v7 := GenS1{F0: 551}
                if true {
                    var v8 float64 = 1.5
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        for fi10 := 0; fi10 < 17; fi10++ {
                            var v11 string = ("a" + ("\n\t" + "hello"))
                            v12 := GenS1{F0: 113}
                            if (!((fi9 + (fi9 >> ((857) & 63))) <= 127)) {
                                if true {
                                    v11 = (v11 + ("\n\t" + "hello"))
                                    v13 := v5[fi10]
                                    _ = v13
                                }
                            }
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        v14 := GenS2{F0: "foo bar", F1: -964}
                        if len(v3) > 0 {
                            v3[0] = fi9
                        }
                        _ = fi9
                        _ = v14
                    }
                }
                var v15 int16 = -32768
                _ = v15
                _ = v7
            } else {
                if ((true && (((float64((-2.25 * ((float32(((((1.5 - 1e10) - float64((0.1 - -2.25))) + float64(float32(0.1))) - 1e10)) * 0.1) * (0.1 - (float32(1e10) + -2.25))))) > float64(3.14159)) && true) || true)) && ("foo bar" == "a")) {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            if len(v3) > 0 {
                                v3[0] = fi17
                            }
                            _ = fi17
                        }
                        var v18 uint64 = 886
                        _ = fi16
                        _ = v18
                    }
                    var v19 float32 = 1.5
                    var v20 bool = true
                    _ = v19
                    _ = v20
                } else {
                    var v21 int8 = 127
                    _ = v21
                }
            }
        } else {
            if true {
                v22 := GenS1{F0: 65535}
                var v23 int8 = 127
                fmt.Println(v23)
                _ = v22
                _ = v23
            }
        }
    } else {
        switch "" {
        case "foo bar":
            v24 := GenS2{F0: ("a" + "foo bar"), F1: -9223372036854775808}
            _ = v24
        case "hello":
            var v25 int = 705
            _ = v25
        default:
            var v26 uint64 = 219
            _ = v26
        }
        var v27 float32 = (-2.25 / float32((((1.5 * (((0.1 + 1e10) / float64(0.1)) * (1.5 * 3.14159))) - ((3.14159 - 3.14159) - ((0.0 / 1.5) / 1e10))) - (0.0 / 1.5))))
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := GenS2{F0: "\n\t", F1: -1}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if ((622 != 37) && false) {
            v3 = append(v3, (fi28 % ((fi28) | 1)))
        } else {
            v30 := GenS2{F0: ("foo bar" + ("\n\t" + "")), F1: 1}
            var v31 int8 = -103
            _ = v30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    if (-838 >= 753) {
        switch "hello" {
        case "a":
            var v32 bool = false
            _ = v32
        case "hello":
            for fi33 := 0; fi33 < 5; fi33++ {
                var v34 float64 = ((3.14159 + (1.5 / 1.5)) * (0.1 * 3.14159))
                _ = fi33
                _ = v34
            }
        }
        var v35 uint = 212
        var v36 int64 = -203
        _ = v35
        _ = v36
    }
    var v37 uint8 = 130
    var v38 string = ("" + "foo bar")
    v38 = v38
    switch -2147483648 {
    case 127:
        v37 -= v37
    }
    fmt.Println(v37, v37)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
