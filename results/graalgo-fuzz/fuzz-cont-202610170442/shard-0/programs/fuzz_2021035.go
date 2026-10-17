package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int64 }

func main() {
    var v1 int8 = 127
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    var v4 int = 0
    var v5 uint16 = 32767
    _ = len(v3)
    fmt.Println(v1, v4)
    v1 = -34
    v4 -= v4
    v6 := []int{(^(v4)), int(v5)}
    v4 = (v4 & v4)
    for fi7 := 0; fi7 < 3; fi7++ {
        switch (("\n\t" + (("a" + "\n\t") + "a")) + ("" + ("hello" + "a"))) {
        case "hello":
            switch (v4 % (((fi7 & fi7)) | 1)) {
            case 879:
                if (false && (1 > 654)) {
                    var v8 int64 = 9223372036854775807
                    if len(v6) > 0 {
                        v6[0] = (v4 / ((fi7) | 1))
                    }
                    v9 := v3[("a" + "\n\t")]
                    _ = v8
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        v2 -= 4294967295
                        if len(v6) > 0 {
                            v6[0] = v4
                        }
                        _ = fi10
                    }
                    var v11 uint64 = 2
                    _ = v11
                }
            case 2147483647:
                if ((true && ((v1 | 100) > -1)) || (v2 != (^(v2)))) {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        switch ((("foo bar" + (("" + ("hello" + (("" + "") + "a"))) + "")) + "") + (("foo bar" + "") + "\n\t")) {
                        case "foo bar":
                            v5 += 1
                        case "a":
                            _ = len(v3)
                        default:
                            v2 = v2
                        }
                        v13 := GenS1{F0: 4294967295, F1: false, F2: 42}
                        _ = fi12
                        _ = v13
                    }
                } else {
                    if (false || (true || false)) {
                        v1 = (-(v1))
                        v4 += v4
                    }
                }
            case -1:
                v14 := GenS1{F0: 65535, F1: true, F2: -9223372036854775808}
                _ = v14
            }
        default:
            var v15 bool = false
            _ = v15
        }
        fmt.Println(v1, v4)
        var v16 float64 = -2.25
        _ = fi7
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            v4 -= v4
            if (18446744073709551615 > 42) {
                if false {
                    _ = len(v6)
                    v4 -= (fi17 >> ((uint(v5)) & 63))
                    v4 = int(v5)
                }
            }
            var v19 int8 = (v1 / (((v1 / ((int8(v5)) | 1))) | 1))
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 int32 = -1
    v5 += (v5 << ((v2) & 63))
    v21 := GenS1{F0: 65535, F1: (true && false), F2: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
