package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 2
    v3 := make(map[int]uint32)
    var v4 uint16 = 284
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int32 = -295
        v4 = (v4 / (((v4 % ((32767) | 1))) | 1))
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 50, F1: 0}
    if false {
        if true {
            v8 := GenS1{F0: v4, F1: -971}
            v4 = 433
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint16 = 42
            v4 = 32767
            _ = fi9
            _ = v10
        }
        if ((v2 < (v2 % (((v2 | 4294967295)) | 1))) && (true && true)) {
            var v11 uint = 1
            v2 = (v2 * v2)
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(v11)
                fmt.Println(v11, v2, v11)
                _ = fi12
            }
            _ = v11
        } else {
            v13 := GenS2{F0: 0, F1: v4}
            _ = v13
        }
        _ = len(v3)
    }
    if (-983 == 1) {
        v2 = 4294967295
    } else {
        switch "foo bar" {
        case "a":
            switch int(v2) {
            case -2147483648:
                delete(v1, ("a" + (("\n\t" + ("foo bar" + ("foo bar" + (("\n\t" + "\n\t") + "")))) + "hello")))
            }
        case "foo bar":
            v2 = v2
        }
        v14 := GenS1{F0: (^(v4)), F1: -968}
        v4 += 0
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS1{F0: (v4 % (((v4 ^ 595)) | 1)), F1: 1}
        v4 = v4
        switch ("\n\t" + "") {
        case "foo bar":
            switch (fi15 << ((4294967295) & 63)) {
            case -1:
                for fi17 := 0; fi17 < 14; fi17++ {
                    v18 := GenS2{F0: -1, F1: (v4 % ((v4) | 1))}
                    if false {
                        var v19 int64 = -9223372036854775808
                        var v20 uint32 = 42
                        var v21 int64 = v19
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int8 = 100
                        _ = v22
                    }
                    var v23 uint64 = v2
                    _ = fi17
                    _ = v18
                    _ = v23
                }
            case -170:
                for fi24 := 0; fi24 < 2; fi24++ {
                    v25 := []int{(fi15 * (fi15 & 127)), fi15, fi15}
                    var v26 int8 = -20
                    v2 = (v2 * 30)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
        }
        _ = fi15
        _ = v16
    }
    v27 := GenS1{F0: 65535, F1: -1}
    v4 += v4
    v2 = (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
