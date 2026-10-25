package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    var v1 float32 = 0.0
    var v2 uint64 = 42
    v3 := []string{("" + (("\n\t" + ("hello" + "\n\t")) + ("foo bar" + ("\n\t" + "\n\t")))), "hello"}
    var v4 int = 640
    var v5 uint = 507
    v4 += v4
    v6 := GenS1{F0: (127 != (-(v4))), F1: ""}
    switch ("hello" + "\n\t") {
    case "hello":
        var v7 uint32 = 2
        _ = v7
    case "a":
        var v8 uint32 = 2147483647
        _ = v8
    case "\n\t":
        v9 := GenS1{F0: false, F1: ((("hello" + (("hello" + ("foo bar" + "hello")) + "")) + "a") + ("\n\t" + ("hello" + "")))}
        _ = v9
    }
    fmt.Println(v4, v2, v2)
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 int = v4
        var v12 float32 = v1
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v4, v5, v2)
    switch (v4 - 9223372036854775807) {
    case -9223372036854775808:
        switch (v4 % ((v4) | 1)) {
        case 976:
            var v13 int8 = 85
            _ = v13
        case -841:
            if len(v3) > 0 {
                v3[0] = ""
            }
        case 127:
            var v14 uint8 = 129
            _ = v14
        }
    case 32768:
        v15 := GenS1{F0: true, F1: "hello"}
        _ = v15
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v5 = (v5 + v5)
    v4 -= int(v5)
    for fi16 := 0; fi16 < 9; fi16++ {
        if ((!(-2.25 >= 0.0)) || (!false)) {
            if false {
                var v17 int = (v4 >> (((v5 - (^(v5)))) & 63))
                _ = v17
            } else {
                var v18 int = fi16
                fmt.Println(v5, fi16, v1)
                switch "" {
                case "":
                    v4 = (v18 << (((^(v5))) & 63))
                }
                _ = v18
            }
            if true {
                if true {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v5 += 458
                        fmt.Println(fi16, fi19, v5)
                        _ = fi19
                    }
                }
                var v20 bool = ("foo bar" > ("" + "foo bar"))
                var v21 uint8 = 242
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := GenS2{F0: 42, F1: 818}
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 4; fi24++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = len(v3)
                    var v25 uint64 = (v2 & 538)
                    _ = fi24
                    _ = v25
                }
            }
            var v26 int8 = -1
            _ = v26
        } else {
            v4 = 1000
            for fi27 := 0; fi27 < 19; fi27++ {
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                fmt.Println(v4)
                var v28 float64 = float64((0.1 + v1))
                _ = fi27
                _ = v28
            }
        }
        _ = fi16
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        if (v2 == 2) {
            v1 -= 0.0
        }
        _ = fi29
    }
    var v30 bool = (!true)
    v3 = append(v3, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
