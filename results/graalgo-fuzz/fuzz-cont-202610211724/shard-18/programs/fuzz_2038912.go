package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint; F1 int32; F2 int16 }
type GenS4 struct { F0 string; F1 int16; F2 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 42
    var v3 uint = 0
    var v4 string = "hello"
    var v5 int32 = -1
    if false {
        if ((963 != -2147483648) || (127 < 137)) {
            v3 -= (^(v3))
            v6 := []int{-1}
            v3 = v3
            _ = v6
        } else {
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 int8 = 79
                switch (fi7 << ((v3) & 63)) {
                case 1:
                    v9 := []int{fi7}
                    _ = v9
                }
                v10 := GenS1{F0: (v2 / ((uint64(v3)) | 1)), F1: v2, F2: -1}
                _ = fi7
                _ = v10
                _ = v8
            }
            var v11 float32 = (1e10 - ((3.14159 * float32((0.1 * float64(1.5)))) * 0.0))
            switch "" {
            case "a":
                v12 := GenS2{F0: 2, F1: -9223372036854775808}
                _ = v12
            }
            _ = v11
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 int64 = -580
            v15 := GenS2{F0: 65535, F1: 248}
            v5 = (v5 * v5)
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: 2, F1: -511}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 string = "\n\t"
            if false {
                for fi19 := 0; fi19 < 16; fi19++ {
                    v3 = v3
                    var v20 uint32 = 42
                    var v21 string = v4
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 7; fi22++ {
                    var v23 uint = 255
                    var v24 float32 = 3.14159
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                var v25 string = "a"
                _ = v25
            } else {
                v4 = (v4 + ((v18 + (v4 + "foo bar")) + v18))
                v26 := GenS2{F0: 523, F1: 128}
                _ = v26
            }
            v27 := GenS3{F0: v3, F1: v5, F2: -1}
            _ = fi17
            _ = v18
            _ = v27
        }
    }
    v5 = v5
    fmt.Println(v2)
    var v28 int16 = 32767
    v29 := GenS3{F0: (v3 | v3), F1: v5, F2: v28}
    var v30 int8 = 49
    if false {
        v5 -= v5
        v4 = v4
        v30 += (v30 >> ((v3) & 63))
    }
    if (1e10 < 1.5) {
        v31 := GenS2{F0: 262, F1: 4294967295}
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 uint16 = 65535
            if false {
                fmt.Println(v33, v33, fi32)
            }
            v34 := GenS2{F0: 521, F1: 9223372036854775807}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v31
    } else {
        var v35 float64 = -2.25
        _ = len(v1)
        _ = v35
    }
    var v36 float32 = 1.5
    for fi37 := 0; fi37 < 5; fi37++ {
        fmt.Println(v4)
        for fi38 := 0; fi38 < 20; fi38++ {
            switch "foo bar" {
            case "foo bar":
                var v39 int8 = -128
                _ = v39
            case "a":
                _ = len(v1)
            }
            v40 := GenS2{F0: 21, F1: (^(fi37))}
            _ = fi38
            _ = v40
        }
        switch (-(fi37)) {
        case -466:
            if (v3 >= v3) {
                switch v4 {
                case "":
                    fmt.Println(fi37)
                }
                fmt.Println(v36, v30)
            }
        case -99:
            var v41 uint8 = 186
            _ = v41
        }
        _ = fi37
    }
    fmt.Println(v4, v3)
    if (3.14159 > float64(v36)) {
        var v42 uint = v3
        _ = v42
    } else {
        v30 = v30
        var v43 uint = (v3 >> ((v3) & 63))
        if false {
            v44 := GenS3{F0: 836, F1: v5, F2: (v28 >> ((1) & 63))}
            for fi45 := 0; fi45 < 17; fi45++ {
                var v46 uint16 = 653
                v47 := GenS4{F0: "", F1: v28, F2: ((v4 + v4) + "hello")}
                v48 := v1[42]
                _ = fi45
                _ = v46
                _ = v47
                _ = v48
            }
            _ = v44
        } else {
            switch ("foo bar" + v4) {
            case "a":
                for fi49 := 0; fi49 < 18; fi49++ {
                    v50 := GenS3{F0: v43, F1: (v5 % (((v5 & (v5 ^ (v5 % (((^(v5))) | 1))))) | 1)), F2: v28}
                    _ = fi49
                    _ = v50
                }
            case "hello":
                v43 += (v43 ^ (v43 % (((v43 << (((v43 % ((42) | 1))) & 63))) | 1)))
            }
        }
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
}
