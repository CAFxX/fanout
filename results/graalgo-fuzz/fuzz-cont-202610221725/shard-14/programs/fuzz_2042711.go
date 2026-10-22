package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int8; F2 int16 }
type GenS5 struct { F0 int8; F1 uint; F2 int64 }
type GenS6 struct { F0 int8; F1 int16; F2 int32 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 int16; F1 int8; F2 int }
type GenS9 struct { F0 int32; F1 int; F2 int64 }
type GenS10 struct { F0 uint8; F1 uint32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS15 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]uint32)
    v2 = (-(v2))
    v2 = 138
    v2 -= v2
    v2 = (^(v2))
    v2 = -1
    if (-2147483648 < 42) {
        var v4 float64 = 1e10
        for fi5 := 0; fi5 < 2; fi5++ {
            for fi6 := 0; fi6 < 6; fi6++ {
                v4 = 3.14159
                var v7 int8 = int8(v2)
                var v8 uint16 = 578
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi5
        }
        if ((v2 << ((255) & 63)) < v2) {
            switch 531 {
            case -9223372036854775808:
                var v9 string = (("" + ("\n\t" + "a")) + "\n\t")
                _ = v9
            case 127:
                var v10 float32 = 3.14159
                _ = v10
            case 42:
                for fi11 := 0; fi11 < 19; fi11++ {
                    switch "foo bar" {
                    case "":
                        v2 -= (v2 | 0)
                    }
                    if ((true || (208 > 85)) || true) {
                        for fi12 := 0; fi12 < 3; fi12++ {
                            if true {
                                if false {
                                    v4 -= (v4 * 1.5)
                                    var v13 int8 = int8(v2)
                                    var v14 string = "\n\t"
                                    var v15 bool = (42 < 255)
                                    fmt.Println(v2, v13, v13)
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                }
                                if (100 != 42) {
                                    var v16 float64 = float64(-2.25)
                                    var v17 uint32 = 772
                                    _ = v16
                                    _ = v17
                                } else {
                                    v2 += v2
                                    switch fi11 {
                                    case 316:
                                        if ((fi12 & -9223372036854775808) >= 127) {
                                            var v18 int64 = 0
                                            _ = v18
                                        }
                                    case 4294967295:
                                        v19 := []int{fi11, (fi12 + fi11)}
                                        _ = v19
                                    case 23:
                                        var v20 int64 = -1
                                        delete(v3, "a")
                                        _ = v20
                                    default:
                                        delete(v3, "foo bar")
                                    }
                                    v21 := []int{-2147483648, (fi11 - 4294967295), 42, fi12}
                                    _ = v21
                                }
                                switch (("" + "\n\t") + ("" + "")) {
                                case "hello":
                                    switch -1 {
                                    case 127:
                                        if (127 != 232) {
                                            switch "\n\t" {
                                            case "hello":
                                                if true {
                                                    switch 846 {
                                                    case 1:
                                                        fmt.Println(v2)
                                                    default:
                                                        if ((42 == 1) && (true && false)) {
                                                            v22 := []int{int(v2)}
                                                            if true {
                                                                v23 := GenS1{F0: 2147483648}
                                                                v4 += v4
                                                                if ((((((1e10 / v4) * float64((0.1 / 3.14159))) * float64(1e10)) < ((float64(((3.14159 - (float32(0.0) / (1e10 + (-2.25 / -2.25)))) - (0.1 * ((1e10 * (3.14159 / (1.5 - ((0.1 / 1.0) / 1.0)))) / 1.5)))) + (v4 + (0.0 / (((v4 * float64(-2.25)) - v4) + v4)))) - v4)) && ((false || (254 <= 0)) || true)) || (!false)) {
                                                                    switch ("\n\t" + "") {
                                                                    case "a":
                                                                        var v24 float64 = float64((3.14159 / ((0.0 - 3.14159) / 3.14159)))
                                                                        _ = v24
                                                                    default:
                                                                        fmt.Println(fi12)
                                                                        var v25 uint = 0
                                                                        fmt.Println(fi11, v4, v4)
                                                                        _ = v25
                                                                    }
                                                                    _ = len(v22)
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v26 := v3["hello"]
                                                                        var v27 uint = 527
                                                                        _ = v26
                                                                        _ = v27
                                                                    }
                                                                } else {
                                                                    var v28 float64 = (v4 - 0.0)
                                                                    var v29 uint = 965
                                                                    if false {
                                                                        var v30 uint64 = 288
                                                                        _ = len(v1)
                                                                        switch ("" + "a") {
                                                                        case "hello":
                                                                            var v31 bool = true
                                                                            _ = v31
                                                                        case "a":
                                                                            var v32 bool = true
                                                                            _ = v32
                                                                        case "\n\t":
                                                                            v33 := GenS1{F0: 2}
                                                                            _ = v33
                                                                        }
                                                                        _ = v30
                                                                    }
                                                                    _ = v28
                                                                    _ = v29
                                                                }
                                                                _ = v23
                                                            } else {
                                                                var v34 bool = false
                                                                if (!((v4 > v4) && v34)) {
                                                                    v34 = (1000 > 589)
                                                                    switch "hello" {
                                                                    case "":
                                                                        v2 += int64(fi12)
                                                                    }
                                                                    v35 := GenS1{F0: 65535}
                                                                    v36 := GenS1{F0: 138}
                                                                    v2 += v2
                                                                    _ = v35
                                                                    _ = v36
                                                                } else {
                                                                    v37 := GenS1{F0: 2147483647}
                                                                    _ = len(v22)
                                                                    fmt.Println(fi12, fi12)
                                                                    _ = v37
                                                                }
                                                                _ = v34
                                                            }
                                                            _ = v22
                                                        }
                                                    }
                                                    fmt.Println(fi11, v4)
                                                } else {
                                                    switch ((("\n\t" + "hello") + (("foo bar" + "\n\t") + "a")) + "foo bar") {
                                                    case "hello":
                                                        fmt.Println(fi11, fi11)
                                                    case "\n\t":
                                                        v1["a"] = 1
                                                    case "foo bar":
                                                        v4 = (v4 + v4)
                                                        if true {
                                                            v2 = 0
                                                            if (255 == 2) {
                                                                var v38 int8 = 0
                                                                fmt.Println(fi11)
                                                                delete(v3, "")
                                                                _ = v38
                                                            } else {
                                                                var v39 float64 = v4
                                                                switch (fi11 >> ((18446744073709551615) & 63)) {
                                                                case 1:
                                                                    fmt.Println(fi12)
                                                                case 127:
                                                                    switch (fi11 % ((-311) | 1)) {
                                                                    case -2147483648:
                                                                        v40 := v1["foo bar"]
                                                                        _ = v40
                                                                    case -918:
                                                                        delete(v1, ("" + (("a" + ("" + (("foo bar" + (("\n\t" + "a") + ("foo bar" + ("\n\t" + "hello")))) + ("\n\t" + ("\n\t" + "hello"))))) + "a")))
                                                                    case -809:
                                                                        switch fi12 {
                                                                        case 785:
                                                                            if false {
                                                                                delete(v1, "")
                                                                                v4 = v4
                                                                            } else {
                                                                                v41 := []int{fi11, 128, fi12}
                                                                                _ = v41
                                                                            }
                                                                        case 127:
                                                                            v2 += (v2 * (v2 ^ (v2 + v2)))
                                                                        default:
                                                                            switch (("hello" + "foo bar") + ("hello" + "hello")) {
                                                                            case "":
                                                                                if (true || true) {
                                                                                    fmt.Println(v2)
                                                                                    switch fi12 {
                                                                                    case -1:
                                                                                        v42 := []int{(fi11 % (((^(fi12))) | 1))}
                                                                                        _ = v42
                                                                                    case 2147483647:
                                                                                        fmt.Println(v39, v39, v2)
                                                                                    case -223:
                                                                                        if (((fi11 & fi11) > 128) || ((int32(fi12) == 1) && (false && false))) {
                                                                                            var v43 string = "hello"
                                                                                            if true {
                                                                                                if true {
                                                                                                    v44 := v3[v43]
                                                                                                    var v45 string = v43
                                                                                                    var v46 int32 = 0
                                                                                                    _ = v44
                                                                                                    _ = v45
                                                                                                    _ = v46
                                                                                                }
                                                                                                switch fi12 {
                                                                                                case 42:
                                                                                                    v47 := GenS1{F0: 387}
                                                                                                    _ = v47
                                                                                                default:
                                                                                                    fmt.Println(fi11)
                                                                                                }
                                                                                            }
                                                                                            v4 = v4
                                                                                            var v48 uint8 = 127
                                                                                            fmt.Println(fi11)
                                                                                            _ = v43
                                                                                            _ = v48
                                                                                        }
                                                                                    }
                                                                                    v49 := GenS1{F0: 891}
                                                                                    if ((v4 >= (0.0 * v39)) && ((-2147483648 >= 0) || true)) {
                                                                                        fmt.Println(fi12, v4, v39)
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            v4 = 3.14159
                                                                                        default:
                                                                                            v50 := GenS2{F0: (v2 * 7), F1: 1, F2: 18446744073709551615}
                                                                                            _ = v50
                                                                                        }
                                                                                        v51 := []int{(fi12 >> ((2) & 63)), fi11, fi11}
                                                                                        switch "" {
                                                                                        case "a":
                                                                                            var v52 int = fi11
                                                                                            _ = v52
                                                                                        case "foo bar":
                                                                                            var v53 uint64 = 472
                                                                                            _ = v53
                                                                                        case "":
                                                                                            v1["foo bar"] = 9223372036854775807
                                                                                        default:
                                                                                            v3["foo bar"] = 2
                                                                                        }
                                                                                        v2 += v2
                                                                                        _ = v51
                                                                                    }
                                                                                    _ = v49
                                                                                }
                                                                            case "a":
                                                                                if true {
                                                                                    v54 := v1["hello"]
                                                                                    if ((-2.25 / (((v39 - (-2.25 + v39)) * (3.14159 + (1.5 + 0.1))) + v4)) < v39) {
                                                                                        switch 1 {
                                                                                        case 2:
                                                                                            var v55 uint32 = 2147483647
                                                                                            _ = v55
                                                                                        case -1:
                                                                                            v56 := GenS3{F0: 1}
                                                                                            _ = v56
                                                                                        }
                                                                                    }
                                                                                    _ = len(v3)
                                                                                    _ = v54
                                                                                } else {
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        var v57 uint64 = 4294967295
                                                                                        _ = v57
                                                                                    }
                                                                                    v58 := GenS3{F0: 404}
                                                                                    fmt.Println(fi12, v4)
                                                                                    _ = v58
                                                                                }
                                                                            case "hello":
                                                                                v59 := []int{fi11}
                                                                                _ = v59
                                                                            default:
                                                                                v60 := GenS4{F0: -254, F1: 127, F2: 553}
                                                                                v39 += (0.0 + 3.14159)
                                                                                var v61 float64 = 0.0
                                                                                _ = v60
                                                                                _ = v61
                                                                            }
                                                                        }
                                                                    }
                                                                case 9223372036854775807:
                                                                    if (!(!(!false))) {
                                                                        var v62 bool = true
                                                                        var v63 string = ""
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                }
                                                                _ = v39
                                                            }
                                                            var v64 string = "hello"
                                                            _ = v64
                                                        } else {
                                                            if true {
                                                                var v65 uint8 = 31
                                                                switch (("a" + "hello") + ("\n\t" + "hello")) {
                                                                case "\n\t":
                                                                    v66 := GenS1{F0: 592}
                                                                    _ = v66
                                                                }
                                                                _ = v65
                                                            }
                                                            var v67 uint16 = 2
                                                            _ = v67
                                                        }
                                                        var v68 uint = 138
                                                        _ = v68
                                                    }
                                                    if ((!true) || false) {
                                                        v69 := GenS2{F0: v2, F1: 4294967295, F2: 385}
                                                        v70 := GenS5{F0: -128, F1: 0, F2: v2}
                                                        v71 := GenS6{F0: 100, F1: -32768, F2: -934}
                                                        _ = v69
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                }
                                            case "":
                                                fmt.Println(fi11, fi11)
                                            case "\n\t":
                                                if (127 > 42) {
                                                    var v72 float64 = (v4 - (1.5 + (3.14159 * v4)))
                                                    _ = v72
                                                }
                                            }
                                            v2 = 726
                                            if (0 != 255) {
                                                var v73 int32 = -95
                                                _ = v73
                                            }
                                        } else {
                                            var v74 uint16 = 2
                                            _ = v74
                                        }
                                    }
                                case "foo bar":
                                    v75 := GenS7{F0: 42, F1: 65535}
                                    _ = v75
                                }
                            }
                            fmt.Println(v4, fi11)
                            v76 := GenS2{F0: (-(v2)), F1: 65535, F2: 100}
                            _ = fi12
                            _ = v76
                        }
                        var v77 int = (fi11 % ((127) | 1))
                        _ = v77
                    } else {
                        v4 = (float64((3.14159 - ((0.0 - 0.0) - ((3.14159 + -2.25) + 3.14159)))) / v4)
                    }
                    if (v4 < v4) {
                        for fi78 := 0; fi78 < 5; fi78++ {
                            var v79 bool = false
                            _ = fi78
                            _ = v79
                        }
                        var v80 uint64 = 2
                        _ = v80
                    }
                    _ = fi11
                }
            default:
                v4 += (((v4 / float64((float32(float64(float32((float64(-2.25) + ((v4 + 3.14159) - 1e10))))) / 0.1))) / ((v4 + v4) * float64((0.1 - float32((0.1 * 3.14159)))))) - v4)
            }
        }
        _ = v4
    } else {
        for fi81 := 0; fi81 < 4; fi81++ {
            if false {
                v82 := GenS7{F0: -32768, F1: 65535}
                v83 := GenS3{F0: 42}
                _ = v82
                _ = v83
            }
            _ = fi81
        }
        v84 := GenS8{F0: 0, F1: 59, F2: 2147483647}
        if ("hello" > "a") {
            v3["\n\t"] = 1
        } else {
            fmt.Println(v2)
        }
        _ = v84
    }
    for fi85 := 0; fi85 < 14; fi85++ {
        for fi86 := 0; fi86 < 2; fi86++ {
            var v87 uint64 = 0
            _ = fi86
            _ = v87
        }
        switch ((("foo bar" + "") + "\n\t") + "foo bar") {
        case "hello":
            v3[("a" + "foo bar")] = 894
        case "\n\t":
            if false {
                delete(v3, "hello")
                var v88 float64 = (float64(1e10) * -2.25)
                for fi89 := 0; fi89 < 16; fi89++ {
                    fmt.Println(v88)
                    switch "foo bar" {
                    case "hello":
                        v88 -= 3.14159
                    case "foo bar":
                        v90 := GenS9{F0: -2147483648, F1: fi89, F2: v2}
                        _ = v90
                    }
                    _ = fi89
                }
                _ = v88
            } else {
                var v91 string = ("" + "")
                for fi92 := 0; fi92 < 8; fi92++ {
                    v93 := GenS7{F0: 0, F1: 32767}
                    var v94 int8 = 42
                    _ = fi92
                    _ = v93
                    _ = v94
                }
                _ = v91
            }
        }
        var v95 int = 32768
        _ = fi85
        _ = v95
    }
    for fi96 := 0; fi96 < 5; fi96++ {
        v2 = (v2 << ((1) & 63))
        _ = fi96
    }
    _ = len(v1)
    v97 := GenS10{F0: 255, F1: 65535}
    var v98 string = ((("\n\t" + ("foo bar" + "a")) + "hello") + "foo bar")
    for fi99 := 0; fi99 < 6; fi99++ {
        v98 = v98
        for fi100 := 0; fi100 < 14; fi100++ {
            v98 = (v98 + v98)
            _ = len(v3)
            if (255 < 2) {
                v2 = v2
                switch fi100 {
                case 815:
                    v101 := []int{fi100, (fi99 % (((fi99 | fi99)) | 1)), 42}
                    _ = v101
                }
                var v102 bool = false
                _ = v102
            }
            if false {
                fmt.Println(v2, fi100)
                v103 := GenS5{F0: 1, F1: 18446744073709551615, F2: 9223372036854775807}
                var v104 int32 = -2147483648
                var v105 bool = ((v2 - (v2 << ((4294967295) & 63))) >= 0)
                _ = v103
                _ = v104
                _ = v105
            } else {
                v2 -= v2
                var v106 int64 = v2
                _ = v106
            }
            v107 := GenS11{F0: 405}
            _ = fi100
            _ = v107
        }
        _ = fi99
    }
    for fi108 := 0; fi108 < 11; fi108++ {
        var v109 int32 = 1
        _ = fi108
        _ = v109
    }
    var v110 int = -1
    if ((-(v2)) == (v2 << ((255) & 63))) {
        switch (v110 % (((^(v110))) | 1)) {
        case -1:
            fmt.Println(v98)
        case 4294967295:
            var v111 uint8 = 42
            _ = v111
        case 32768:
            for fi112 := 0; fi112 < 5; fi112++ {
                var v113 float64 = 0.1
                if (2 <= 4294967295) {
                    switch ("\n\t" + v98) {
                    case "hello":
                        fmt.Println(v98, v110)
                    case "foo bar":
                        for fi114 := 0; fi114 < 4; fi114++ {
                            if (true && (false && (2 < 404))) {
                                v2 -= (v2 ^ v2)
                                fmt.Println(fi114)
                                var v115 float64 = 1e10
                                v116 := []int{(fi114 & (fi112 % ((-741) | 1))), fi112, fi114}
                                _ = v115
                                _ = v116
                            }
                            v117 := []int{(-(v110)), 465}
                            v113 += 0.1
                            _ = fi114
                            _ = v117
                        }
                    default:
                        var v118 uint8 = 187
                        _ = v118
                    }
                    v110 = (^(fi112))
                    v110 -= (fi112 - fi112)
                } else {
                    for fi119 := 0; fi119 < 3; fi119++ {
                        fmt.Println(v110, fi112)
                        if (v98 == v98) {
                            switch 480 {
                            case 32767:
                                if (true || (!((!false) && false))) {
                                    v2 = (v2 * (v2 ^ (v2 | 0)))
                                } else {
                                    fmt.Println(v110)
                                    switch fi112 {
                                    case -2147483648:
                                        var v120 uint64 = 1
                                        _ = v120
                                    case 32767:
                                        if false {
                                            var v121 uint32 = 227
                                            var v122 int = v110
                                            _ = v121
                                            _ = v122
                                        }
                                    }
                                    var v123 float64 = v113
                                    switch v98 {
                                    case "a":
                                        v123 -= v113
                                    case "hello":
                                        var v124 int32 = 2147483647
                                        _ = v124
                                    default:
                                        v123 -= 0.1
                                    }
                                    _ = v123
                                }
                            case 128:
                                if (2 < 184) {
                                    if (32767 < 1) {
                                        if false {
                                            fmt.Println(v2, v110)
                                            v113 = (v113 - (((v113 + -2.25) / -2.25) / v113))
                                            switch fi112 {
                                            case -9223372036854775808:
                                                var v125 uint = 436
                                                _ = v125
                                            default:
                                                switch ((v98 + v98) + v98) {
                                                case "hello":
                                                    _ = len(v1)
                                                case "":
                                                    switch (v110 >> ((18446744073709551615) & 63)) {
                                                    case 2147483647:
                                                        v126 := GenS11{F0: 1}
                                                        _ = v126
                                                    }
                                                case "a":
                                                    switch int(v2) {
                                                    case 421:
                                                        v127 := []int{}
                                                        v110 = fi112
                                                        _ = v127
                                                    case 0:
                                                        fmt.Println(fi112)
                                                    }
                                                }
                                            }
                                            var v128 bool = true
                                            var v129 bool = (32767 >= -94)
                                            _ = v128
                                            _ = v129
                                        } else {
                                            var v130 bool = (!true)
                                            var v131 uint64 = 4294967295
                                            fmt.Println(v2, v131, v98)
                                            _ = v130
                                            _ = v131
                                        }
                                    } else {
                                        var v132 int8 = 100
                                        fmt.Println(v132, v113)
                                        switch (v110 / (((fi119 >> ((546) & 63))) | 1)) {
                                        case -2147483648:
                                            fmt.Println(v113)
                                        }
                                        v2 = (v2 >> ((255) & 63))
                                        var v133 uint64 = 896
                                        v132 += (v132 / ((42) | 1))
                                        delete(v3, v98)
                                        _ = v132
                                        _ = v133
                                    }
                                    v134 := GenS10{F0: 2, F1: 197}
                                    _ = v134
                                } else {
                                    if ((-(v110)) == int(v2)) {
                                        if true {
                                            fmt.Println(v2, v98)
                                            v135 := []int{}
                                            if false {
                                                v98 = "foo bar"
                                                v113 = float64(float32(-2.25))
                                                if false {
                                                    v110 += (-(v110))
                                                    v98 = "\n\t"
                                                }
                                            }
                                            _ = v135
                                        } else {
                                            var v136 uint16 = 67
                                            _ = v136
                                        }
                                        if false {
                                            _ = len(v3)
                                            var v137 float32 = ((1e10 + (-2.25 + 3.14159)) / 3.14159)
                                            _ = v137
                                        } else {
                                            v98 = (v98 + v98)
                                        }
                                        v138 := GenS12{F0: 65535, F1: 0, F2: 1}
                                        var v139 string = "foo bar"
                                        var v140 int16 = -32768
                                        _ = v138
                                        _ = v139
                                        _ = v140
                                    }
                                    if true {
                                        if (((!(!(!true))) || (!false)) && true) {
                                            v141 := []int{-2147483648, fi119, (v110 / ((fi112) | 1))}
                                            switch v98 {
                                            case "\n\t":
                                                if false {
                                                    switch v98 {
                                                    case "":
                                                        v141 = append(v141, 454)
                                                    case "hello":
                                                        if ((fi119 << ((693) & 63)) < 0) {
                                                            v113 += v113
                                                            var v142 int64 = v2
                                                            fmt.Println(v113, fi112)
                                                            _ = v142
                                                        } else {
                                                            v98 = ("a" + (((v98 + ((v98 + ("hello" + (((v98 + "\n\t") + (v98 + v98)) + (v98 + v98)))) + v98)) + "foo bar") + (v98 + "a")))
                                                        }
                                                    default:
                                                        var v143 uint16 = 2
                                                        _ = v143
                                                    }
                                                    delete(v1, v98)
                                                }
                                                var v144 float32 = ((1e10 - (3.14159 / 1e10)) - 1e10)
                                                v1[(v98 + (v98 + v98))] = 9223372036854775807
                                                _ = v144
                                            default:
                                                if (("" >= "") && true) {
                                                    if ((!true) && true) {
                                                        v98 = v98
                                                        if (4294967295 != 119) {
                                                            v145 := GenS13{F0: 42}
                                                            v146 := GenS14{F0: 941, F1: 1, F2: 40}
                                                            _ = v145
                                                            _ = v146
                                                        }
                                                    }
                                                    fmt.Println(fi119)
                                                }
                                            }
                                            v98 = ((v98 + "\n\t") + v98)
                                            _ = v141
                                        }
                                    } else {
                                        if true {
                                            v110 -= v110
                                            if true {
                                                var v147 int64 = -526
                                                _ = v147
                                            } else {
                                                switch (v110 - (fi119 & 0)) {
                                                case 2147483647:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if true {
                                                            v110 -= 0
                                                            var v148 int = fi112
                                                            _ = v148
                                                        }
                                                    default:
                                                        fmt.Println(fi112)
                                                    }
                                                case 32768:
                                                    v2 = int64(fi112)
                                                case 0:
                                                    v149 := []int{}
                                                    _ = v149
                                                }
                                                var v150 uint32 = 0
                                                fmt.Println(v150, fi112, v113)
                                                v98 = v98
                                                v110 -= fi112
                                                _ = v150
                                            }
                                        }
                                        v2 -= (v2 >> ((0) & 63))
                                        var v151 float64 = ((1.5 * v113) * v113)
                                        v152 := []int{int(v2)}
                                        _ = v151
                                        _ = v152
                                    }
                                    delete(v3, v98)
                                }
                            case 4294967295:
                                var v153 uint = 162
                                _ = len(v1)
                                if false {
                                    v154 := []int{(fi119 & (v110 * fi112))}
                                    v98 = (v98 + v98)
                                    if (true || true) {
                                        v153 = (v153 % ((v153) | 1))
                                        v155 := GenS15{F0: 32767, F1: 128}
                                        _ = v155
                                    }
                                    v156 := GenS13{F0: 1}
                                    v113 = v113
                                    _ = v154
                                    _ = v156
                                }
                                _ = v153
                            default:
                                var v157 float32 = float32(0.1)
                                _ = v157
                            }
                            v110 = (fi119 % ((325) | 1))
                        }
                        v98 = v98
                        var v158 string = "foo bar"
                        v113 -= 1.5
                        _ = fi119
                        _ = v158
                    }
                    v110 += fi112
                    fmt.Println(v98)
                }
                _ = fi112
                _ = v113
            }
        }
    } else {
        v159 := v1["foo bar"]
        _ = v159
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v98:", v98)
    _ = v1
    _ = v110
    _ = v2
    _ = v3
    _ = v97
    _ = v98
}
