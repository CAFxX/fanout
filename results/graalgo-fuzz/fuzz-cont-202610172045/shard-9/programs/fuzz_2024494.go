package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []string{"foo bar", ("\n\t" + ("" + (("\n\t" + "hello") + "\n\t"))), "hello"}
    v2 := make(map[string]int)
    v3 := []string{"", "foo bar", "\n\t", ("" + "foo bar")}
    var v4 string = (("\n\t" + "a") + "\n\t")
    v3 = append(v3, "")
    v4 = v4
    var v5 uint16 = 595
    delete(v2, v4)
    v4 = ((((v4 + (v4 + v4)) + v4) + ("a" + v4)) + "\n\t")
    if (65535 < 42) {
        if true {
            for fi6 := 0; fi6 < 6; fi6++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (!((float32((1e10 - (0.1 * 0.1))) * (0.0 - 0.0)) >= (0.1 * float32(0.1)))) {
                    switch ((v4 + ((v4 + v4) + (("a" + v4) + v4))) + "\n\t") {
                    case "":
                        if (false && false) {
                            v4 = (v4 + (v4 + "foo bar"))
                        }
                    case "foo bar":
                        v7 := v2["hello"]
                        _ = v7
                    }
                    var v8 uint64 = 0
                    _ = v8
                }
                if (v4 >= v4) {
                    if true {
                        delete(v2, v4)
                        v9 := GenS1{F0: 0}
                        _ = v9
                    } else {
                        for fi10 := 0; fi10 < 16; fi10++ {
                            switch 128 {
                            case 9223372036854775807:
                                v5 = 988
                            case -2147483648:
                                var v11 float64 = 0.0
                                _ = v11
                            }
                            _ = fi10
                        }
                        if true {
                            if true {
                                v12 := GenS1{F0: -32768}
                                v4 = "a"
                                _ = v12
                            }
                        } else {
                            if (42 >= 0) {
                                v4 = "\n\t"
                                v4 = (v4 + v4)
                            }
                        }
                        v4 = v4
                    }
                    v4 = v4
                }
                _ = fi6
            }
            switch "foo bar" {
            case "hello":
                v2[(("" + "hello") + "\n\t")] = 128
            }
        }
    } else {
        _ = len(v3)
        v3 = append(v3, (v4 + (v4 + v4)))
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            if (!true) {
                v4 = "\n\t"
            } else {
                var v15 float32 = (0.0 / 1e10)
                _ = v15
            }
            var v16 int8 = 100
            _ = fi14
            _ = v16
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            fmt.Println(v5, fi13)
            v4 = ("" + ((v4 + v4) + "a"))
            switch 42 {
            case 9223372036854775807:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v4 = v4
            default:
                var v18 int64 = 1
                _ = v18
            }
            _ = fi17
        }
        _ = fi13
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 bool = true
        _ = fi19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
