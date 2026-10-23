package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int8 }

func main() {
    var v1 bool = true
    var v2 uint16 = 65535
    v3 := make(map[string]uint32)
    v4 := []int{777, 4294967295, 1}
    var v5 int = 0
    var v6 float64 = 3.14159
    v3[(("foo bar" + "hello") + "a")] = 271
    if ((v2 ^ (v2 >> ((4294967295) & 63))) < 32767) {
        var v7 int8 = 1
        v1 = ((^(v5)) <= v5)
        var v8 bool = (42 < -841)
        _ = v7
        _ = v8
    }
    var v9 int8 = 1
    switch (v5 ^ v5) {
    case 2147483647:
        switch (("foo bar" + "\n\t") + ("\n\t" + ("\n\t" + "hello"))) {
        case "\n\t":
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 uint8 = 86
                v6 = v6
                var v12 uint32 = 42
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            if len(v4) > 0 {
                v4[0] = 0
            }
        }
    case -442:
        var v13 string = ((("hello" + ("" + "")) + "a") + ("a" + ((("\n\t" + ("hello" + ((((((("a" + (("\n\t" + ("foo bar" + ("\n\t" + ""))) + (("hello" + "hello") + ("a" + "")))) + (((("foo bar" + "foo bar") + "") + "\n\t") + (("hello" + ("\n\t" + ("foo bar" + ""))) + ("hello" + "hello")))) + "a") + "a") + "foo bar") + "") + ("a" + ("\n\t" + ((((("foo bar" + "foo bar") + ("" + "")) + "") + ("foo bar" + "")) + "foo bar")))))) + ("a" + ("a" + "foo bar"))) + "foo bar")))
        _ = v13
    }
    if v1 {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 int16 = -485
            _ = fi14
            _ = v15
        }
    }
    if (!false) {
        for fi16 := 0; fi16 < 1; fi16++ {
            switch int(v2) {
            case -632:
                v6 += (v6 * 0.1)
            case 128:
                if v1 {
                    fmt.Println(v6, v2)
                    if v1 {
                        v6 += v6
                        _ = len(v4)
                        v17 := v3["hello"]
                        _ = v17
                    }
                }
            }
            var v18 int16 = 0
            _ = fi16
            _ = v18
        }
    }
    v19 := GenS1{F0: v2, F1: 2, F2: v9}
    switch "" {
    case "":
        if v1 {
            var v20 uint64 = 42
            if ((!v1) || (0 <= -1)) {
                v2 -= 0
            } else {
                switch 42 {
                case -275:
                    for fi21 := 0; fi21 < 19; fi21++ {
                        if v1 {
                            v22 := GenS1{F0: 1, F1: 0, F2: 17}
                            _ = v22
                        } else {
                            v23 := GenS1{F0: v2, F1: 0, F2: (v9 >> ((255) & 63))}
                            v20 -= (^(v20))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v23
                        }
                        fmt.Println(v9, v6, v9)
                        if (3.14159 != float32(v6)) {
                            v3["hello"] = 364
                            if ((!(!v1)) && (((!v1) && (!(222 <= 1))) && (!(-1 < -32768)))) {
                                if v1 {
                                    var v24 float32 = 0.0
                                    var v25 string = "foo bar"
                                    _ = v24
                                    _ = v25
                                }
                            } else {
                                v6 -= 1.5
                            }
                        }
                        _ = fi21
                    }
                }
                var v26 string = ("foo bar" + "\n\t")
                _ = v26
            }
            _ = v20
        } else {
            var v27 int16 = 32767
            v27 -= (v27 >> ((319) & 63))
            _ = v27
        }
    case "hello":
        var v28 float32 = 1e10
        _ = v28
    case "a":
        v29 := GenS1{F0: v2, F1: 460, F2: (v9 << ((518) & 63))}
        _ = v29
    }
    v30 := GenS1{F0: (v2 << ((695) & 63)), F1: 2, F2: 63}
    for fi31 := 0; fi31 < 8; fi31++ {
        var v32 uint16 = v2
        if (1 == 0) {
            v4 = append(v4, fi31)
            for fi33 := 0; fi33 < 15; fi33++ {
                if false {
                    switch (-(v5)) {
                    case -15:
                        v34 := GenS1{F0: 0, F1: 4294967295, F2: (v9 | -120)}
                        _ = v34
                    case 32768:
                        var v35 bool = v1
                        v36 := v3["\n\t"]
                        v37 := GenS1{F0: 65535, F1: 905, F2: -1}
                        _ = v35
                        _ = v36
                        _ = v37
                    default:
                        v2 = 65535
                    }
                }
                fmt.Println(v32)
                _ = fi33
            }
            v6 = v6
        }
        _ = fi31
        _ = v32
    }
    v6 = v6
    if (!true) {
        v3["a"] = 0
    }
    var v38 int8 = 1
    var v39 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
