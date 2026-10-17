package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 string = "a"
    var v4 float64 = 0.0
    v5 := []int{127, 0, -583, 32768}
    var v6 string = ((("" + v3) + v3) + ("\n\t" + v3))
    var v7 bool = (42 >= 253)
    v8 := GenS1{F0: v7}
    v5 = append(v5, 32767)
    delete(v1, 0)
    v3 = ""
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint64 = 18446744073709551615
            v7 = v7
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 uint8 = 127
    var v13 int16 = -1
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 bool = false
        fmt.Println(v3, v12)
        for fi16 := 0; fi16 < 6; fi16++ {
            if v7 {
                v13 -= v13
                switch fi16 {
                case -9223372036854775808:
                    v3 = "a"
                }
                var v17 float32 = (1.5 / float32(1e10))
                _ = v17
            } else {
                switch (("\n\t" + (("\n\t" + (("foo bar" + v6) + ("hello" + v3))) + "hello")) + "hello") {
                case "":
                    var v18 int8 = 127
                    _ = v18
                case "hello":
                    if (false || (!true)) {
                        var v19 uint32 = 65535
                        if ((-9223372036854775808 <= 0) || v15) {
                            var v20 uint16 = 65535
                            fmt.Println(v20, v12, v19)
                            var v21 string = "\n\t"
                            _ = v20
                            _ = v21
                        } else {
                            if v15 {
                                var v22 uint8 = (v12 << ((18446744073709551615) & 63))
                                _ = v22
                            }
                            v15 = (4294967295 >= 2)
                        }
                        _ = v19
                    } else {
                        fmt.Println(v6)
                        v7 = ((0 > 4294967295) && v15)
                    }
                case "\n\t":
                    var v23 string = (((v6 + v3) + (v6 + "foo bar")) + "\n\t")
                    _ = v23
                }
            }
            switch ((v3 + v6) + "\n\t") {
            case "":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            switch (-(fi14)) {
            case -1:
                v15 = (!v15)
            case 32767:
                v24 := GenS1{F0: (1 == 1)}
                _ = v24
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v7, v13, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
