package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 string; F2 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    var v3 bool = true
    v4 := make(map[int]bool)
    v5 := []string{"foo bar", ("\n\t" + "\n\t"), ""}
    var v6 string = "\n\t"
    v6 = "foo bar"
    v7 := GenS1{F0: -1}
    var v8 bool = (!v3)
    if (-36 < -9223372036854775808) {
        var v9 int = 127
        if (762 > 2147483648) {
            if (v6 == (v6 + v6)) {
                if true {
                    var v10 uint64 = 211
                    _ = v10
                }
            } else {
                v11 := GenS1{F0: -1}
                if v8 {
                    if (18446744073709551615 >= 52) {
                        v12 := GenS2{F0: "\n\t", F1: 775, F2: 530}
                        fmt.Println(v8, v6)
                        v3 = false
                        _ = v12
                    }
                }
                _ = v11
            }
        }
        _ = v9
    } else {
        if (!(true || v3)) {
            v3 = (221 >= 4294967295)
        } else {
            switch (v6 + "\n\t") {
            case "":
                for fi13 := 0; fi13 < 6; fi13++ {
                    if (fi13 >= fi13) {
                        v14 := GenS1{F0: -154}
                        v6 = ("\n\t" + v6)
                        v15 := GenS1{F0: 1}
                        _ = v14
                        _ = v15
                    }
                    _ = fi13
                }
            }
            v8 = (42 == -32768)
        }
        v6 = (v6 + "hello")
    }
    var v16 uint8 = 1
    if v8 {
        for fi17 := 0; fi17 < 1; fi17++ {
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 int8 = -1
                _ = fi18
                _ = v19
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi17
        }
    } else {
        v20 := GenS2{F0: (("a" + ("foo bar" + (v6 + ("a" + ("foo bar" + ("foo bar" + v6)))))) + ((v6 + ("\n\t" + v6)) + v6)), F1: 0, F2: 1}
        v21 := GenS3{F0: 42}
        _ = v20
        _ = v21
    }
    v22 := GenS4{F0: 2147483648}
    if (!false) {
        v3 = v8
        v23 := GenS5{F0: 9223372036854775807}
        v6 = v6
        _ = v23
    } else {
        v24 := GenS6{F0: (("" + v6) + "a"), F1: "\n\t", F2: int16(v16)}
        var v25 string = "foo bar"
        fmt.Println(v8, v25)
        _ = v24
        _ = v25
    }
    if len(v5) > 0 {
        v5[0] = "a"
    }
    if v3 {
        v6 = v6
        v26 := v1[""]
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
