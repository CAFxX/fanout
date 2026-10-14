package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }

func main() {
    var v1 int16 = -558
    var v2 int16 = (v1 ^ (-(v1)))
    v3 := make(map[string]bool)
    fmt.Println(v2)
    for fi4 := 0; fi4 < 10; fi4++ {
        v2 -= 1000
        _ = fi4
    }
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 bool = true
        v2 = -32768
        v7 := GenS1{F0: 42, F1: 2}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{127, 4294967295, -2147483648, 2147483647}
    var v9 int64 = -1
    var v10 int = 32767
    if (((!(!true)) && (!true)) || (!((1.5 == 1e10) && true))) {
        fmt.Println(v1, v9)
        v2 = v2
        if ((v10 << ((2) & 63)) >= (-(v10))) {
            for fi11 := 0; fi11 < 18; fi11++ {
                fmt.Println(fi11, v10)
                _ = fi11
            }
        }
    }
    if true {
        var v12 string = ("" + ("\n\t" + ""))
        delete(v3, v12)
        fmt.Println(v12, v9, v10)
        _ = v12
    } else {
        if (true && ((v10 >= v10) || (!false))) {
            v9 -= (v9 - v9)
            v2 += v2
            if (false && (127 >= 42)) {
                if ((false && ("\n\t" == "\n\t")) || (false || (42 < 933))) {
                    var v13 int8 = -96
                    _ = v13
                }
                fmt.Println(v1, v1)
                switch (("hello" + (("foo bar" + "hello") + ((("foo bar" + "a") + "\n\t") + ("a" + "\n\t")))) + "\n\t") {
                case "foo bar":
                    if false {
                        fmt.Println(v10, v2)
                        _ = len(v3)
                        v10 += -9223372036854775808
                    }
                case "\n\t":
                    var v14 float32 = ((0.0 - 3.14159) * -2.25)
                    _ = v14
                }
            } else {
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 float64 = float64(0.0)
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 12; fi17++ {
                    if false {
                        var v18 uint8 = 2
                        v18 = (^(v18))
                        _ = v18
                    }
                    if false {
                        for fi19 := 0; fi19 < 14; fi19++ {
                            fmt.Println(v1, v10, fi17)
                            switch -946 {
                            case 42:
                                var v20 uint32 = 2147483648
                                _ = v20
                            case -1:
                                v8 = append(v8, (fi17 - 4294967295))
                            default:
                                v10 = v10
                            }
                            _ = fi19
                        }
                        if (127 > 146) {
                            var v21 int = fi17
                            if (((!(65535 <= 42)) || (-33 != (v1 * v1))) && (!(0.0 >= float32(1e10)))) {
                                for fi22 := 0; fi22 < 11; fi22++ {
                                    if true {
                                        fmt.Println(v2, v10)
                                        delete(v3, ("hello" + "foo bar"))
                                    } else {
                                        v9 = (^(v9))
                                    }
                                    if (false || (980 >= 65535)) {
                                        var v23 uint16 = 1
                                        var v24 int8 = -128
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = fi22
                                }
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    v26 := GenS1{F0: 4294967295, F1: 414}
                                    var v27 string = "a"
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                            }
                            _ = v21
                        }
                        for fi28 := 0; fi28 < 9; fi28++ {
                            v29 := GenS1{F0: 0, F1: 255}
                            _ = len(v3)
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        v8 = append(v8, (fi17 >> ((171) & 63)))
                    }
                    v3[("foo bar" + ("\n\t" + "a"))] = (42 >= 1)
                    _ = fi17
                }
            }
        }
    }
    var v30 int8 = 0
    switch "" {
    case "\n\t":
        v9 = (v9 % ((v9) | 1))
    case "a":
        for fi31 := 0; fi31 < 9; fi31++ {
            v32 := GenS1{F0: 0, F1: 32767}
            v3[""] = false
            _ = fi31
            _ = v32
        }
    case "hello":
        if false {
            for fi33 := 0; fi33 < 13; fi33++ {
                v10 -= -2147483648
                v2 = v1
                _ = fi33
            }
        }
    }
    var v34 int = 485
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v8
    _ = v9
}
