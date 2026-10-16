package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]uint32)
    v3 := []string{"\n\t"}
    v4 := make(map[int]string)
    v5 := []int{-942}
    for fi6 := 0; fi6 < 2; fi6++ {
        v1 = (v1 | (v1 & 1))
        delete(v2, (("hello" + "hello") + "a"))
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint = 426
            var v9 int32 = -4
            v10 := GenS1{F0: (v8 % (((^(v8))) | 1)), F1: 824, F2: 1000}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint16 = 32767
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint32 = 836
            if true {
                if (false && ((-(fi11)) != 32768)) {
                    if true {
                        var v15 string = ("foo bar" + ("" + ("foo bar" + "hello")))
                        v16 := GenS1{F0: 0, F1: 683, F2: -252}
                        _ = v15
                        _ = v16
                    }
                    v1 += (v1 - 1)
                }
                var v17 int8 = -96
                switch "\n\t" {
                case "":
                    v1 = (v1 | 1)
                case "hello":
                    v18 := GenS1{F0: 371, F1: 2, F2: 401}
                    _ = v18
                }
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    if ((!(91 > 0)) && false) {
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    }
    v1 += 0
    v1 = 0
    var v19 string = "a"
    v19 = "foo bar"
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 float32 = (0.1 / -2.25)
        if true {
            v5 = append(v5, 158)
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 int16 = -32768
                fmt.Println(v19, v19)
                v23 = -735
                _ = fi22
                _ = v23
            }
        } else {
            v24 := GenS1{F0: 255, F1: 42, F2: 42}
            _ = v24
        }
        fmt.Println(v21, v19, fi20)
        _ = fi20
        _ = v21
    }
    v19 = (((v19 + "foo bar") + v19) + ("\n\t" + ("" + v19)))
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 int64 = 423
        fmt.Println(v26)
        _ = fi25
        _ = v26
    }
    var v27 uint64 = 0
    switch 2147483647 {
    case 32767:
        v3 = append(v3, (((v19 + v19) + "a") + ("hello" + v19)))
    }
    if (!(true && (65535 >= 65535))) {
        v28 := GenS1{F0: 2, F1: (v27 - (v27 + v27)), F2: 769}
        _ = v28
    }
    if ((42 != 1) || false) {
        var v29 int32 = -2147483648
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
